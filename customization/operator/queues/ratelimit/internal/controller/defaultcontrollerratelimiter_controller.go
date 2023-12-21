/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sync"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// DefaultControllerRateLimiterReconciler reconciles a DefaultControllerRateLimiter object
type DefaultControllerRateLimiterReconciler struct {
	client.Client
	Scheme   *runtime.Scheme
	PodsMap  map[ctrl.Request]int
	Mutex    *sync.RWMutex
	requeues int
}

////+kubebuilder:rbac:groups="",resources=pod,verbs=get;list;watch
//+kubebuilder:rbac:groups=ratelimit.shalabi.com,resources=defaultcontrollerratelimiters,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=ratelimit.shalabi.com,resources=defaultcontrollerratelimiters/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=ratelimit.shalabi.com,resources=defaultcontrollerratelimiters/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the DefaultControllerRateLimiter object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.3/pkg/reconcile

func (r *DefaultControllerRateLimiterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := log.FromContext(ctx)

	r.Mutex.Lock()
	v := r.PodsMap[req]
	r.PodsMap[req] = v + 1
	r.requeues++
	r.Mutex.Unlock()

	l.Info("erorrs", "count", r.PodsMap[req], "time", time.Now())
	l.Info("requeues number", "count", r.requeues)
	return ctrl.Result{}, fmt.Errorf("some error")
}

// SetupWithManager sets up the controller with the Manager.
func (r *DefaultControllerRateLimiterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1.Pod{}).
		WithOptions(controller.Options{
			MaxConcurrentReconciles: 10,
			//RateLimiter: workqueue.DefaultControllerRateLimiter(),
			RateLimiter: &workqueue.BucketRateLimiter{Limiter: rate.NewLimiter(rate.Limit(10), 100)},
		}).
		Complete(r)
}
