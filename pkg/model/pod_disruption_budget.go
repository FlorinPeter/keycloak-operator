package model

import (
	"github.com/keycloak/keycloak-operator/pkg/apis/keycloak/v1alpha1"
	policyv1 "k8s.io/api/policy/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func PodDisruptionBudget(cr *v1alpha1.Keycloak) *policyv1.PodDisruptionBudget {
	return &policyv1.PodDisruptionBudget{
		ObjectMeta: v1.ObjectMeta{
			Name:      ApplicationName,
			Namespace: cr.Namespace,
			Labels: map[string]string{
				"app": ApplicationName,
			},
		},
		Spec: policyv1.PodDisruptionBudgetSpec{
			MaxUnavailable: &intstr.IntOrString{IntVal: MaxUnavailableNumberOfPods},
			Selector: &v1.LabelSelector{
				MatchLabels: map[string]string{"component": KeycloakDeploymentComponent},
			},
		},
	}
}

func PodDisruptionBudgetReconciled(cr *v1alpha1.Keycloak, currentState *policyv1.PodDisruptionBudget) *policyv1.PodDisruptionBudget {
	reconciled := currentState.DeepCopy()
	reconciled.Spec.MaxUnavailable = &intstr.IntOrString{IntVal: MaxUnavailableNumberOfPods}
	reconciled.Spec.Selector = &v1.LabelSelector{
		MatchLabels: map[string]string{"component": KeycloakDeploymentComponent},
	}
	return reconciled
}

func PodDisruptionBudgetSelector(cr *v1alpha1.Keycloak) client.ObjectKey {
	return client.ObjectKey{
		Name:      ApplicationName,
		Namespace: cr.Namespace,
	}
}
