/*
Copyright 2020 Michael Bridgen <mikeb@squaremobius.net>

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

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	imagev1alpha1 "github.com/squaremo/image-automation-controller/api/v1alpha1"
)

// ImageUpdateAutomationReconciler reconciles a ImageUpdateAutomation object
type ImageUpdateAutomationReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=image.fluxcd.io,resources=imageupdateautomations,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=image.fluxcd.io,resources=imageupdateautomations/status,verbs=get;update;patch

func (r *ImageUpdateAutomationReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("imageupdateautomation", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *ImageUpdateAutomationReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&imagev1alpha1.ImageUpdateAutomation{}).
		Complete(r)
}