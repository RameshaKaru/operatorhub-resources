package certificateauthority

import (
	"io/ioutil"
	"path/filepath"

	corev1 "k8s.io/api/core/v1"
)

const (
	// TrustedCaConfigMapName is the name of the Marketplace ConfigMap that stores Certificate Authority bundle.
	TrustedCaConfigMapName = "marketplace-trusted-ca"

	// trustedCaMountPath is the path to the directory where the Certificate Authority volume should be mounted.
	trustedCaMountPath = "/etc/pki/ca-trust/extracted/pem/"

	// caBundleKey is the key in the ConfigMap that stores Certificate Authoritie bundle.
	caBundleKey = "ca-bundle.crt"

	// caBundlePath is the path where we will mount the Certificate Authorities bundle.
	caBundlePath = "tls-ca-bundle.pem"
)

// MountCaConfigMap adds a Volume and VolumeMount for the Certificate Authority ConfigMap on
// the given PodTemplateSpec.
func MountCaConfigMap(template *corev1.PodTemplateSpec) {
	// Create and add the Volume to the PodTemplateSpec.
	template.Spec.Volumes = []corev1.Volume{
		corev1.Volume{
			Name: TrustedCaConfigMapName,
			VolumeSource: corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: TrustedCaConfigMapName,
					},
					Items: []corev1.KeyToPath{
						corev1.KeyToPath{
							Key:  caBundleKey,
							Path: caBundlePath,
						},
					},
				},
			},
		},
	}

	// Create and add the VolumeMount to the first container in the PodTemplateSpec.
	template.Spec.Containers[0].VolumeMounts = []corev1.VolumeMount{
		corev1.VolumeMount{
			Name:      TrustedCaConfigMapName,
			MountPath: trustedCaMountPath,
		},
	}
}

// getCaOnDisk returns the contents of the Certificate Authority bundle on disk as a byte
// array or returns the error encountered when attempting to do so.
func getCaOnDisk() ([]byte, error) {
	return ioutil.ReadFile(filepath.Join(trustedCaMountPath, caBundlePath))
}
