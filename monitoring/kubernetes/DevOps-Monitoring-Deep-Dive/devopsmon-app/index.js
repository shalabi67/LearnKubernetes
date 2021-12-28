// Require and call Express
var express = require('express');
var bodyParser = require('body-parser');
var app = express();

var responseTime = require('response-time');
const prom = require('prom-client');
const collectDefaultMetrics = prom.collectDefaultMetrics;
collectDefaultMetrics({ prefix: 'forethought' });

// Prometheus metric definitions
 const todocounter = new prom.Counter({
   name: 'forethought_number_of_todos_total',
   help: 'The number of items added to the to-do list, total'
 });
 const todogauge = new prom.Gauge ({
   name: 'forethought_current_todos',
   help: 'Amount of incomplete tasks'
 });
 const tasksumm = new prom.Summary ({
    name: 'forethought_requests_summ',
    help: 'Latency in percentiles',
  });
  const taskhisto = new prom.Histogram ({
    name: 'forethought_requests_hist',
    help: 'Latency in history form',
    // this can be defined later after we run the application for a while and we have an idea about out bucket sizes
    buckets: [0.1, 0.25, 0.5, 1, 2.5, 5, 10]
  });

app.use(bodyParser.urlencoded({ extended: true }));
app.set('view engine', 'ejs');

// use css
app.use(express.static("public"));

// placeholder tasks
var task = [];
var complete = [];


app.use(responseTime(function (req, res, time) {
   tasksumm.observe(time);
   taskhisto.observe(time);
 }));

// add a task
app.post("/addtask", function(req, res) {
  var newTask = req.body.newtask;
  task.push(newTask);
  res.redirect("/");
  todocounter.inc();
  todogauge.inc();
});

// remove a task
app.post("/removetask", function(req, res) {
  var completeTask = req.body.check;
  if (typeof completeTask === "string") {
    complete.push(completeTask);
    task.splice(task.indexOf(completeTask), 1);
    todogauge.dec();
  }
  else if (typeof completeTask === "object") {
    for (var i = 0; i < completeTask.length; i++) {
      complete.push(completeTask[i]);
      task.splice(task.indexOf(completeTask[i]), 1);
      todogauge.dec();
    }
  }
  res.redirect("/");
});

// get website files
app.get("/", function (req, res) {
  res.render("index", { task: task, complete: complete });
});

 app.get('/metrics', async function (req, res) {
   res.set('Content-Type', prom.register.contentType);
   res.end(await prom.register.metrics());
 });

// listen for connections
app.listen(8080, function() {
  console.log('Testing app listening on port 8080')
});
