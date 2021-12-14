// Copyright YEAR VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	appv1alpha1 "github.com/vmware-tanzu/tanzu-framework/apis/app/v1alpha1"
	configv1alpha1 "github.com/vmware-tanzu/tanzu-framework/apis/config/v1alpha1"
	"github.com/vmware-tanzu/tanzu-framework/pkg/v1/sdk/features/featuregate"
)

// LookupReconciler reconciles a Lookup object
type LookupReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=app.tanzu.vmware.com,resources=lookups,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=app.tanzu.vmware.com,resources=lookups/status,verbs=get;update;patch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Lookup object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.6.4/pkg/reconcile
func (r *LookupReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("lookup", req.NamespacedName)

	activated, err := featuregate.FeatureActivatedInNamespace(ctx, r.Client, "tkg-system", "tinycache")
	if err != nil {
		return ctrl.Result{}, err
	}
	if !activated {
		log.Info("Feature 'tinycache' deactivated")
		log.Info("Skipping reconcile")
		return ctrl.Result{}, nil
	}

	log.Info("Using 'tinycache' feature")
	log.Info("Successfully reconciled")
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *LookupReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appv1alpha1.Lookup{}).
		Watches(&source.Kind{Type: &configv1alpha1.FeatureGate{}}, handler.EnqueueRequestsFromMapFunc(r.featureGatesToLookupList)).
		Complete(r)
}

func (r *LookupReconciler) featureGatesToLookupList(o client.Object) []reconcile.Request {
	var requests []reconcile.Request

	lookupList := &appv1alpha1.LookupList{}
	if err := r.Client.List(context.Background(), lookupList); err != nil {
		r.Log.Error(err, "failed to list lookupList in event handler")
		return requests
	}

	for i := range lookupList.Items {
		requests = append(requests, reconcile.Request{
			NamespacedName: types.NamespacedName{
				Namespace: lookupList.Items[i].Namespace,
				Name:      lookupList.Items[i].Name,
			},
		})
	}

	return requests
}
