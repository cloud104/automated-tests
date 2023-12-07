package k8s

import (
	corev1 "k8s.io/api/core/v1"
)

func IsPodStatusConditionTrue(conditions []corev1.PodCondition, conditionType corev1.PodConditionType) bool {
	return IsPodStatusConditionPresentAndEqual(conditions, conditionType, corev1.ConditionTrue)
}

func IsPodStatusConditionPresentAndEqual(conditions []corev1.PodCondition, conditionType corev1.PodConditionType, status corev1.ConditionStatus) bool {
	for _, condition := range conditions {
		if condition.Type == conditionType {
			return condition.Status == status
		}
	}

	return false
}
