package kserve

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	servingv1 "knative.dev/serving/pkg/apis/serving/v1"
	servingclient "knative.dev/serving/pkg/client/clientset/versioned"
)

// Client represents a KServe client
type Client struct {
	k8sClient     *kubernetes.Clientset
	servingClient *servingclient.Clientset
	namespace     string
}

// NewClient creates a new KServe client
func NewClient(namespace string) (*Client, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get in-cluster config: %v", err)
	}

	k8sClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create kubernetes client: %v", err)
	}

	servingClient, err := servingclient.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create serving client: %v", err)
	}

	return &Client{
		k8sClient:     k8sClient,
		servingClient: servingClient,
		namespace:     namespace,
	}, nil
}

// DeployModel deploys a model using KServe
func (c *Client) DeployModel(ctx context.Context, name string, modelURI string, framework string) error {
	service := &servingv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: c.namespace,
		},
		Spec: servingv1.ServiceSpec{
			ConfigurationSpec: servingv1.ConfigurationSpec{
				Template: servingv1.RevisionTemplateSpec{
					Spec: servingv1.RevisionSpec{
						PodSpec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Name:  "model",
									Image: fmt.Sprintf("kserve/%s:latest", framework),
									Env: []corev1.EnvVar{
										{
											Name:  "MODEL_URI",
											Value: modelURI,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	_, err := c.servingClient.ServingV1().Services(c.namespace).Create(ctx, service, metav1.CreateOptions{})
	return err
}

// GetModelStatus gets the status of a deployed model
func (c *Client) GetModelStatus(ctx context.Context, name string) (*servingv1.Service, error) {
	return c.servingClient.ServingV1().Services(c.namespace).Get(ctx, name, metav1.GetOptions{})
}

// ListModels lists all deployed models
func (c *Client) ListModels(ctx context.Context) (*servingv1.ServiceList, error) {
	return c.servingClient.ServingV1().Services(c.namespace).List(ctx, metav1.ListOptions{})
}

// DeleteModel deletes a deployed model
func (c *Client) DeleteModel(ctx context.Context, name string) error {
	return c.servingClient.ServingV1().Services(c.namespace).Delete(ctx, name, metav1.DeleteOptions{})
}
