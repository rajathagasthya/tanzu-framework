// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"sigs.k8s.io/cluster-api/cmd/clusterctl/cmd"

	"github.com/vmware-tanzu-private/tkg-cli/pkg/constants"
	"github.com/vmware-tanzu-private/tkg-cli/pkg/log"
	"github.com/vmware-tanzu-private/tkg-cli/pkg/tkgctl"
)

type initRegionOptions struct {
	clusterConfigFile           string
	plan                        string
	ui                          bool
	clusterName                 string
	useExistingCluster          bool
	coreProvider                string
	bootstrapProvider           string
	infrastructureProvider      string
	controlPlaneProvider        string
	targetNamespace             string
	watchingNamespace           string
	timeout                     time.Duration
	size                        string
	controlPlaneSize            string
	workerSize                  string
	tmcRegistrationURL          string
	disableYTT                  bool
	ceipOptIn                   string
	cniType                     string
	featureFlags                map[string]string
	enableTKGSOnVsphere7        bool
	deployTKGonVsphere7         bool
	bind                        string
	browser                     string
	vsphereControlPlaneEndpoint string
	unattended                  bool
}

var iro = &initRegionOptions{}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a Tanzu Kubernetes Grid management cluster",
	Long: cmd.LongDesc(`
			Create a Tanzu Kubernetes Grid management cluster including initializing it with Cluster API components appropriate for the target infrastructure.
		`),

	Example: `
    # Create a management cluster on AWS infrastructure, initializing it with
    # components required to create workload clusters through it on the same infrastructure
    # by bootstrapping through a self-provisioned bootstrap cluster.
    tanzu management-cluster create --file ~/clusterconfigs/aws-mc-1.yaml
    # Launch an interactive UI to configure the settings necessary to create a
    # management cluster
    tanzu management-cluster create --ui
    # Create a management cluster on vSphere infrastructure by using an existing
    # bootstrapper cluster. The current kube context should point to that
    # of the existing bootstrap cluster.
    tanzu management-cluster create --use-existing-bootstrap-cluster --file vsphere-mc-1.yaml`,

	RunE: func(cmd *cobra.Command, args []string) error {
		return runInit()
	},
}

func init() {
	createCmd.Flags().StringVarP(&iro.clusterConfigFile, "file", "f", "", "Configuration file from which to create a management cluster")

	createCmd.Flags().BoolVarP(&iro.ui, "ui", "u", false, "Launch interactive management cluster provisioning UI")
	createCmd.Flags().StringVarP(&iro.bind, "bind", "b", "127.0.0.1:8080", "Specify the IP and port to bind the Kickstart UI against (e.g. 127.0.0.1:8080).")
	createCmd.Flags().StringVarP(&iro.browser, "browser", "", "", "Specify the browser to open the Kickstart UI on. Use 'none' for no browser. Defaults to OS default browser. Supported: ['chrome', 'firefox', 'safari', 'ie', 'edge', 'none']")

	createCmd.Flags().BoolVarP(&iro.unattended, "yes", "y", false, "Create management cluster without asking for confirmation")

	createCmd.Flags().BoolVarP(&iro.useExistingCluster, "use-existing-bootstrap-cluster", "e", false, "Use an existing bootstrap cluster to deploy the management cluster")
	createCmd.Flags().DurationVarP(&iro.timeout, "timeout", "t", constants.DefaultLongRunningOperationTimeout, "Time duration to wait for an operation before timeout. Timeout duration in hours(h)/minutes(m)/seconds(s) units or as some combination of them (e.g. 2h, 30m, 2h30m10s)")

	createCmd.Flags().StringVarP(&iro.infrastructureProvider, "infrastructure", "i", "", "Infrastructure to deploy the management cluster on ['aws', 'vsphere', 'azure']")
	createCmd.Flags().MarkHidden("infrastructure")

	createCmd.Flags().StringVarP(&iro.plan, "plan", "p", "", "Cluster plan to use to deploy the management cluster")
	createCmd.Flags().MarkHidden("plan")

	createCmd.Flags().StringVarP(&iro.clusterName, "name", "", "", "Name of the management cluster. One will be generated if not provided")
	createCmd.Flags().MarkHidden("name")

	createCmd.Flags().StringVarP(&iro.size, "size", "", "", "Specify size for all nodes including control plane and worker nodes. It can be overridden by --controlplane-size and --worker-size options. (See [+])")
	createCmd.Flags().MarkHidden("size")

	createCmd.Flags().StringVarP(&iro.controlPlaneSize, "controlplane-size", "", "", "Specify size for the control plane node. (See [+])")
	createCmd.Flags().MarkHidden("controlplane-size")

	createCmd.Flags().StringVarP(&iro.workerSize, "worker-size", "", "", "Specify size of the worker node. (See [+])")
	createCmd.Flags().MarkHidden("worker-size")

	createCmd.Flags().StringVarP(&iro.ceipOptIn, "ceip-participation", "", "", "Specify if this management cluster should participate in VMware CEIP. (See [*])")
	createCmd.Flags().MarkHidden("ceip-participation")

	createCmd.Flags().BoolVarP(&iro.deployTKGonVsphere7, "deploy-tkg-on-vSphere7", "", false, "Deploy TKG Management cluster on vSphere 7.0 without prompt")
	createCmd.Flags().MarkHidden("deploy-tkg-on-vSphere7")

	createCmd.Flags().BoolVarP(&iro.enableTKGSOnVsphere7, "enable-tkgs-on-vSphere7", "", false, "Enable TKGS on vSphere 7.0 without prompt")
	createCmd.Flags().MarkHidden("enable-tkgs-on-vSphere7")

	createCmd.Flags().StringVarP(&iro.vsphereControlPlaneEndpoint, "vsphere-controlplane-endpoint", "", "", "Virtual IP address or FQDN for the cluster's control plane nodes")
	createCmd.Flags().MarkHidden("vsphere-controlplane-endpoint")

	createCmd.Flags().StringVarP(&iro.tmcRegistrationURL, "tmc-registration-url", "", "", "URL to download the yml which has configuration related to resources to be deployed on the management cluster for it to register with Tanzu Mission Control")
	createCmd.Flags().MarkHidden("tmc-registration-url")

	// Hidden flags, mostly for development and testing

	createCmd.Flags().StringVarP(&iro.targetNamespace, "target-namespace", "", "", "The target namespace where the providers should be deployed. If not specified, each provider will be installed in a provider's default namespace")
	createCmd.Flags().MarkHidden("target-namespace") //nolint

	createCmd.Flags().StringVarP(&iro.cniType, "cni", "", "", "Specify the CNI provider the management cluster should use ['antrea' (default), 'calico', 'none'].")
	createCmd.Flags().MarkHidden("cni") //nolint

	createCmd.Flags().StringToStringVarP(&iro.featureFlags, "feature-flags", "", nil, "Enable and disable hidden features in the form 'feature1=true,feature2=false'")
	createCmd.Flags().MarkHidden("feature-flags") //nolint

	createCmd.Flags().SetNormalizeFunc(aliasNormalizeFunc)
}

func aliasNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	switch name {
	case "vsphere-controlplane-endpoint-ip":
		name = "vsphere-controlplane-endpoint"
	}
	return pflag.NormalizedName(name)
}

func runInit() error {
	tkgClient, err := newTKGCtlClient()
	if err != nil {
		return err
	}

	options := tkgctl.InitRegionOptions{
		ClusterConfigFile:           iro.clusterConfigFile,
		Plan:                        iro.plan,
		UI:                          iro.ui,
		ClusterName:                 iro.clusterName,
		UseExistingCluster:          iro.useExistingCluster,
		CoreProvider:                iro.coreProvider,
		BootstrapProvider:           iro.bootstrapProvider,
		InfrastructureProvider:      iro.infrastructureProvider,
		ControlPlaneProvider:        iro.controlPlaneProvider,
		Namespace:                   iro.targetNamespace,
		WatchingNamespace:           iro.watchingNamespace,
		Size:                        iro.size,
		ControlPlaneSize:            iro.controlPlaneSize,
		WorkerSize:                  iro.workerSize,
		TmcRegistrationURL:          iro.tmcRegistrationURL,
		CeipOptIn:                   iro.ceipOptIn,
		CniType:                     iro.cniType,
		FeatureFlags:                iro.featureFlags,
		EnableTKGSOnVsphere7:        iro.enableTKGSOnVsphere7,
		DeployTKGonVsphere7:         iro.deployTKGonVsphere7,
		Bind:                        iro.bind,
		Browser:                     iro.browser,
		VsphereControlPlaneEndpoint: iro.vsphereControlPlaneEndpoint,
		SkipPrompt:                  iro.unattended,
		Timeout:                     iro.timeout,
	}

	err = tkgClient.Init(options)
	if err == nil {
		log.Info("\nYou can now create your first workload cluster by running the following:\n\n")
		log.Info("  tanzu cluster create [name] -f [file]\n\n")
	}

	return err
}
