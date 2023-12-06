package k8s

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

func IsDeploymentStatusConditionTrue(conditions []appsv1.DeploymentCondition, conditionType appsv1.DeploymentConditionType) bool {
	return IsDeploymentStatusConditionPresentAndEqual(conditions, conditionType, corev1.ConditionTrue)
}

func IsDeploymentStatusConditionPresentAndEqual(conditions []appsv1.DeploymentCondition, conditionType appsv1.DeploymentConditionType, status corev1.ConditionStatus) bool {
	for _, condition := range conditions {
		if condition.Type == conditionType {
			return condition.Status == status
		}
	}

	return false
}
