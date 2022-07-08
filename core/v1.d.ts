import { Condition, LabelSelector, ListMeta, MicroTime, ObjectMeta, Time } from "../meta/v1";
import { Quantity } from "../api/resource";
/**
 * Represents a Persistent Disk resource in AWS.
 *
 * An AWS EBS disk must exist before mounting to a container. The disk must also be in the same AWS zone as the kubelet. An AWS EBS disk can only be mounted as read/write once. AWS EBS volumes support ownership management and SELinux relabeling.
 */
export type AWSElasticBlockStoreVolumeSource = {
    /**
     * fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
     */
    fsType?: string;
    /**
     * partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as "1". Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty).
     */
    partition?: number;
    /**
     * readOnly value true will force the readOnly setting in VolumeMounts. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
     */
    readOnly?: boolean;
    /**
     * volumeID is unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
     */
    volumeID: string;
}
/**
 * Affinity is a group of affinity scheduling rules.
 */
export type Affinity = {
    /**
     * Describes node affinity scheduling rules for the pod.
     */
    nodeAffinity?: NodeAffinity;
    /**
     * Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).
     */
    podAffinity?: PodAffinity;
    /**
     * Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).
     */
    podAntiAffinity?: PodAntiAffinity;
}
/**
 * AttachedVolume describes a volume attached to a node
 */
export type AttachedVolume = {
    /**
     * DevicePath represents the device path where the volume should be available
     */
    devicePath: string;
    /**
     * Name of the attached volume
     */
    name: string;
}
/**
 * AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
 */
export type AzureDiskVolumeSource = {
    /**
     * cachingMode is the Host Caching mode: None, Read Only, Read Write.
     */
    cachingMode?: string;
    /**
     * diskName is the Name of the data disk in the blob storage
     */
    diskName: string;
    /**
     * diskURI is the URI of data disk in the blob storage
     */
    diskURI: string;
    /**
     * fsType is Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
     */
    fsType?: string;
    /**
     * kind expected values are Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared
     */
    kind?: string;
    /**
     * readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
     */
    readOnly?: boolean;
}
/**
 * AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
 */
export type AzureFilePersistentVolumeSource = {
    /**
     * readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
     */
    readOnly?: boolean;
    /**
     * secretName is the name of secret that contains Azure Storage Account Name and Key
     */
    secretName: string;
    /**
     * secretNamespace is the namespace of the secret that contains Azure Storage Account Name and Key default is the same as the Pod
     */
    secretNamespace?: string;
    /**
     * shareName is the azure Share Name
     */
    shareName: string;
}
/**
 * AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
 */
export type AzureFileVolumeSource = {
    /**
     * readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
     */
    readOnly?: boolean;
    /**
     * secretName is the  name of secret that contains Azure Storage Account Name and Key
     */
    secretName: string;
    /**
     * shareName is the azure share Name
     */
    shareName: string;
}
/**
 * Binding ties one object to another; for example, a pod is bound to a node by a scheduler. Deprecated in 1.7, please use the bindings subresource of pods instead.
 */
export type Binding = {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "v1";
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "Binding";
    /**
     * Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     */
    metadata?: ObjectMeta;
    /**
     * The target object that you want to bind to the standard object.
     */
    target: ObjectReference;
}
/**
 * Represents storage that is managed by an external CSI volume driver (Beta feature)
 */
export type CSIPersistentVolumeSource = {
    /**
     * controllerExpandSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI ControllerExpandVolume call. This is an beta field and requires enabling ExpandCSIVolumes feature gate. This field is optional, and may be empty if no secret is required. If the secret object contains more than one secret, all secrets are passed.
     */
    controllerExpandSecretRef?: SecretReference;
    /**
     * controllerPublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI ControllerPublishVolume and ControllerUnpublishVolume calls. This field is optional, and may be empty if no secret is required. If the secret object contains more than one secret, all secrets are passed.
     */
    controllerPublishSecretRef?: SecretReference;
    /**
     * driver is the name of the driver to use for this volume. Required.
     */
    driver: string;
    /**
     * fsType to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs".
     */
    fsType?: string;
    /**
     * nodeExpandSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodeExpandVolume call. This is an alpha field and requires enabling CSINodeExpandSecret feature gate. This field is optional, may be omitted if no secret is required. If the secret object contains more than one secret, all secrets are passed.
     */
    nodeExpandSecretRef?: SecretReference;
    /**
     * nodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and may be empty if no secret is required. If the secret object contains more than one secret, all secrets are passed.
     */
    nodePublishSecretRef?: SecretReference;
    /**
     * nodeStageSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodeStageVolume and NodeStageVolume and NodeUnstageVolume calls. This field is optional, and may be empty if no secret is required. If the secret object contains more than one secret, all secrets are passed.
     */
    nodeStageSecretRef?: SecretReference;
    /**
     * readOnly value to pass to ControllerPublishVolumeRequest. Defaults to false (read/write).
     */
    readOnly?: boolean;
    /**
     * volumeAttributes of the volume to publish.
     */
    volumeAttributes?: {
        [name: string]: string;
    };
    /**
     * volumeHandle is the unique volume name returned by the CSI volume plugin’s CreateVolume to refer to the volume on all subsequent calls. Required.
     */
    volumeHandle: string;
}
/**
 * Represents a source location of a volume to mount, managed by an external CSI driver
 */
export type CSIVolumeSource = {
    /**
     * driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster.
     */
    driver: string;
    /**
     * fsType to mount. Ex. "ext4", "xfs", "ntfs". If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply.
     */
    fsType?: string;
    /**
     * nodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and  may be empty if no secret is required. If the secret object contains more than one secret, all secret references are passed.
     */
    nodePublishSecretRef?: LocalObjectReference;
    /**
     * readOnly specifies a read-only configuration for the volume. Defaults to false (read/write).
     */
    readOnly?: boolean;
    /**
     * volumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's documentation for supported values.
     */
    volumeAttributes?: {
        [name: string]: string;
    };
}
/**
 * Adds and removes POSIX capabilities from running containers.
 */
export type Capabilities = {
    /**
     * Added capabilities
     */
    add?: Array<string>;
    /**
     * Removed capabilities
     */
    drop?: Array<string>;
}
/**
 * Represents a Ceph Filesystem mount that lasts the lifetime of a pod Cephfs volumes do not support ownership management or SELinux relabeling.
 */
export type CephFSPersistentVolumeSource = {
    /**
     * monitors is Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
     */
    monitors: Array<string>;
    /**
     * path is Optional: Used as the mounted root, rather than the full Ceph tree, default is /
     */
    path?: string;
    /**
     * readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
     */
    readOnly?: boolean;
    /**
     * secretFile is Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
     */
    secretFile?: string;
    /**
     * secretRef is Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
     */
    secretRef?: SecretReference;
    /**
     * user is Optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
     */
    user?: string;
}
/**
 * Represents a Ceph Filesystem mount that lasts the lifetime of a pod Cephfs volumes do not support ownership management or SELinux relabeling.
 */
export type CephFSVolumeSource = {
    /**
     * monitors is Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
     */
    monitors: Array<string>;
    /**
     * path is Optional: Used as the mounted root, rather than the full Ceph tree, default is /
     */
    path?: string;
    /**
     * readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
     */
    readOnly?: boolean;
    /**
     * secretFile is Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
     */
    secretFile?: string;
    /**
     * secretRef is Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
     */
    secretRef?: LocalObjectReference;
    /**
     * user is optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
     */
    user?: string;
}
/**
 * Represents a cinder volume resource in Openstack. A Cinder volume must exist before mounting to a container. The volume must also be in the same region as the kubelet. Cinder volumes support ownership management and SELinux relabeling.
 */
export type CinderPersistentVolumeSource = {
    /**
     * fsType Filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
     */
    fsType?: string;
    /**
     * readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
     */
    readOnly?: boolean;
    /**
     * secretRef is Optional: points to a secret object containing parameters used to connect to OpenStack.
     */
    secretRef?: SecretReference;
    /**
     * volumeID used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
     */
    volumeID: string;
}
/**
 * Represents a cinder volume resource in Openstack. A Cinder volume must exist before mounting to a container. The volume must also be in the same region as the kubelet. Cinder volumes support ownership management and SELinux relabeling.
 */
export type CinderVolumeSource = {
    /**
     * fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
     */
    fsType?: string;
    /**
     * readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
     */
    readOnly?: boolean;
    /**
     * secretRef is optional: points to a secret object containing parameters used to connect to OpenStack.
     */
    secretRef?: LocalObjectReference;
    /**
     * volumeID used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
     */
    volumeID: string;
}
/**
 * ClientIPConfig represents the configurations of Client IP based session affinity.
 */
export type ClientIPConfig = {
    /**
     * timeoutSeconds specifies the seconds of ClientIP type session sticky time. The value must be >0 && <=86400(for 1 day) if ServiceAffinity == "ClientIP". Default value is 10800(for 3 hours).
     */
    timeoutSeconds?: number;
}
/**
 * Information about the condition of a component.
 */
export type ComponentCondition = {
    /**
     * Condition error code for a component. For example, a health check error code.
     */
    error?: string;
    /**
     * Message about the condition for a component. For example, information about a health check.
     */
    message?: string;
    /**
     * Status of the condition for a component. Valid values for "Healthy": "True", "False", or "Unknown".
     */
    status: string;
    /**
     * Type of condition for a component. Valid value: "Healthy"
     */
    type: string;
}
/**
 * ComponentStatus (and ComponentStatusList) holds the cluster validation info. Deprecated: This API is deprecated in v1.19+
 */
export type ComponentStatus = {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "v1";
    /**
     * List of component conditions observed
     */
    conditions?: Array<ComponentCondition>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "ComponentStatus";
    /**
     * Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     */
    metadata?: ObjectMeta;
}
/**
 * Status of all the conditions for the component as a list of ComponentStatus objects. Deprecated: This API is deprecated in v1.19+
 */
export type ComponentStatusList = {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "v1";
    /**
     * List of ComponentStatus objects.
     */
    items: Array<ComponentStatus>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "ComponentStatusList";
    /**
     * Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    metadata?: ListMeta;
}
/**
 * ConfigMap holds configuration data for pods to consume.
 */
export type ConfigMap = {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "v1";
    /**
     * BinaryData contains the binary data. Each key must consist of alphanumeric characters, '-', '_' or '.'. BinaryData can contain byte sequences that are not in the UTF-8 range. The keys stored in BinaryData must not overlap with the ones in the Data field, this is enforced during validation process. Using this field will require 1.10+ apiserver and kubelet.
     */
    binaryData?: {
        [name: string]: string;
    };
    /**
     * Data contains the configuration data. Each key must consist of alphanumeric characters, '-', '_' or '.'. Values with non-UTF-8 byte sequences must use the BinaryData field. The keys stored in Data must not overlap with the keys in the BinaryData field, this is enforced during validation process.
     */
    data?: {
        [name: string]: string;
    };
    /**
     * Immutable, if set to true, ensures that data stored in the ConfigMap cannot be updated (only object metadata can be modified). If not set to true, the field can be modified at any time. Defaulted to nil.
     */
    immutable?: boolean;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "ConfigMap";
    /**
     * Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     */
    metadata?: ObjectMeta;
}
/**
 * ConfigMapEnvSource selects a ConfigMap to populate the environment variables with.
 *
 * The contents of the target ConfigMap's Data field will represent the key-value pairs as environment variables.
 */
export type ConfigMapEnvSource = {
    /**
     * Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
     */
    name?: string;
    /**
     * Specify whether the ConfigMap must be defined
     */
    optional?: boolean;
}
/**
 * Selects a key from a ConfigMap.
 */
export type ConfigMapKeySelector = {
    /**
     * The key to select.
     */
    key: string;
    /**
     * Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
     */
    name?: string;
    /**
     * Specify whether the ConfigMap or its key must be defined
     */
    optional?: boolean;
}
/**
 * ConfigMapList is a resource containing a list of ConfigMap objects.
 */
export type ConfigMapList = {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "v1";
    /**
     * Items is the list of ConfigMaps.
     */
    items: Array<ConfigMap>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "ConfigMapList";
    /**
     * More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     */
    metadata?: ListMeta;
}
/**
 * ConfigMapNodeConfigSource contains the information to reference a ConfigMap as a config source for the Node. This API is deprecated since 1.22: https://git.k8s.io/enhancements/keps/sig-node/281-dynamic-kubelet-configuration
 */
export type ConfigMapNodeConfigSource = {
    /**
     * KubeletConfigKey declares which key of the referenced ConfigMap corresponds to the KubeletConfiguration structure This field is required in all cases.
     */
    kubeletConfigKey: string;
    /**
     * Name is the metadata.name of the referenced ConfigMap. This field is required in all cases.
     */
    name: string;
    /**
     * Namespace is the metadata.namespace of the referenced ConfigMap. This field is required in all cases.
     */
    namespace: string;
    /**
     * ResourceVersion is the metadata.ResourceVersion of the referenced ConfigMap. This field is forbidden in Node.Spec, and required in Node.Status.
     */
    resourceVersion?: string;
    /**
     * UID is the metadata.UID of the referenced ConfigMap. This field is forbidden in Node.Spec, and required in Node.Status.
     */
    uid?: string;
}
/**
 * Adapts a ConfigMap into a projected volume.
 *
 * The contents of the target ConfigMap's Data field will be presented in a projected volume as files using the keys in the Data field as the file names, unless the items element is populated with specific mappings of keys to paths. Note that this is identical to a configmap volume source without the default mode.
 */
export type ConfigMapProjection = {
    /**
     * items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.
     */
    items?: Array<KeyToPath>;
    /**
     * Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
     */
    name?: string;
    /**
     * optional specify whether the ConfigMap or its keys must be defined
     */
    optional?: boolean;
}
/**
 * Adapts a ConfigMap into a volume.
 *
 * The contents of the target ConfigMap's Data field will be presented in a volume as files using the keys in the Data field as the file names, unless the items element is populated with specific mappings of keys to paths. ConfigMap volumes support ownership management and SELinux relabeling.
 */
export type ConfigMapVolumeSource = {
    /**
     * defaultMode is optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
     */
    defaultMode?: number;
    /**
     * items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.
     */
    items?: Array<KeyToPath>;
    /**
     * Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
     */
    name?: string;
    /**
     * optional specify whether the ConfigMap or its keys must be defined
     */
    optional?: boolean;
}
/**
 * A single application container that you want to run within a pod.
 */
export type Container = {
    /**
     * Arguments to the entrypoint. The container image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)" will produce the string literal "$(VAR_NAME)". Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
     */
    args?: Array<string>;
    /**
     * Entrypoint array. Not executed within a shell. The container image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)" will produce the string literal "$(VAR_NAME)". Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
     */
    command?: Array<string>;
    /**
     * List of environment variables to set in the container. Cannot be updated.
     */
    env?: Array<EnvVar>;
    /**
     * List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated.
     */
    envFrom?: Array<EnvFromSource>;
    /**
     * Container image name. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets.
     */
    image?: string;
    /**
     * Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images
     *
     *
     */
    imagePullPolicy?: string;
    /**
     * Actions that the management system should take in response to container lifecycle events. Cannot be updated.
     */
    lifecycle?: Lifecycle;
    /**
     * Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
     */
    livenessProbe?: Probe;
    /**
     * Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.
     */
    name: string;
    /**
     * List of ports to expose from the container. Exposing a port here gives the system additional information about the network connections a container uses, but is primarily informational. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default "0.0.0.0" address inside a container will be accessible from the network. Cannot be updated.
     */
    ports?: Array<ContainerPort>;
    /**
     * Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
     */
    readinessProbe?: Probe;
    /**
     * Compute Resources required by this container. Cannot be updated. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
     */
    resources?: ResourceRequirements;
    /**
     * SecurityContext defines the security options the container should be run with. If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext. More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
     */
    securityContext?: SecurityContext;
    /**
     * StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this completes successfully. If this probe fails, the Pod will be restarted, just as if the livenessProbe failed. This can be used to provide different probe parameters at the beginning of a Pod's lifecycle, when it might take a long time to load data or warm a cache, than during steady-state operation. This cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
     */
    startupProbe?: Probe;
    /**
     * Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.
     */
    stdin?: boolean;
    /**
     * Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false
     */
    stdinOnce?: boolean;
    /**
     * Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated.
     */
    terminationMessagePath?: string;
    /**
     * Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.
     *
     *
     */
    terminationMessagePolicy?: string;
    /**
     * Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false.
     */
    tty?: boolean;
    /**
     * volumeDevices is the list of block devices to be used by the container.
     */
    volumeDevices?: Array<VolumeDevice>;
    /**
     * Pod volumes to mount into the container's filesystem. Cannot be updated.
     */
    volumeMounts?: Array<VolumeMount>;
    /**
     * Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.
     */
    workingDir?: string;
}
/**
 * Describe a container image
 */
export type ContainerImage = {
    /**
     * Names by which this image is known. e.g. ["kubernetes.example/hyperkube:v1.0.7", "cloud-vendor.registry.example/cloud-vendor/hyperkube:v1.0.7"]
     */
    names?: Array<string>;
    /**
     * The size of the image in bytes.
     */
    sizeBytes?: number;
}
/**
 * ContainerPort represents a network port in a single container.
 */
export type ContainerPort = {
    /**
     * Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.
     */
    containerPort: number;
    /**
     * What host IP to bind the external port to.
     */
    hostIP?: string;
    /**
     * Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.
     */
    hostPort?: number;
    /**
     * If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services.
     */
    name?: string;
    /**
     * Protocol for port. Must be UDP, TCP, or SCTP. Defaults to "TCP".
     *
     *
     */
    protocol?: string;
}
/**
 * ContainerState holds a possible state of container. Only one of its members may be specified. If none of them is specified, the default one is ContainerStateWaiting.
 */
export type ContainerState = {
    /**
     * Details about a running container
     */
    running?: ContainerStateRunning;
    /**
     * Details about a terminated container
     */
    terminated?: ContainerStateTerminated;
    /**
     * Details about a waiting container
     */
    waiting?: ContainerStateWaiting;
}
/**
 * ContainerStateRunning is a running state of a container.
 */
export type ContainerStateRunning = {
    /**
     * Time at which the container was last (re-)started
     */
    startedAt?: Time;
}
/**
 * ContainerStateTerminated is a terminated state of a container.
 */
export type ContainerStateTerminated = {
    /**
     * Container's ID in the format '<type>://<container_id>'
     */
    containerID?: string;
    /**
     * Exit status from the last termination of the container
     */
    exitCode: number;
    /**
     * Time at which the container last terminated
     */
    finishedAt?: Time;
    /**
     * Message regarding the last termination of the container
     */
    message?: string;
    /**
     * (brief) reason from the last termination of the container
     */
    reason?: string;
    /**
     * Signal from the last termination of the container
     */
    signal?: number;
    /**
     * Time at which previous execution of the container started
     */
    startedAt?: Time;
}
/**
 * ContainerStateWaiting is a waiting state of a container.
 */
export type ContainerStateWaiting = {
    /**
     * Message regarding why the container is not yet running.
     */
    message?: string;
    /**
     * (brief) reason the container is not yet running.
     */
    reason?: string;
}
/**
 * ContainerStatus contains details for the current status of this container.
 */
export type ContainerStatus = {
    /**
     * Container's ID in the format '<type>://<container_id>'.
     */
    containerID?: string;
    /**
     * The image the container is running. More info: https://kubernetes.io/docs/concepts/containers/images.
     */
    image: string;
    /**
     * ImageID of the container's image.
     */
    imageID: string;
    /**
     * Details about the container's last termination condition.
     */
    lastState?: ContainerState;
    /**
     * This must be a DNS_LABEL. Each container in a pod must have a unique name. Cannot be updated.
     */
    name: string;
    /**
     * Specifies whether the container has passed its readiness probe.
     */
    ready: boolean;
    /**
     * The number of times the container has been restarted.
     */
    restartCount: number;
    /**
     * Specifies whether the container has passed its startup probe. Initialized as false, becomes true after startupProbe is considered successful. Resets to false when the container is restarted, or if kubelet loses state temporarily. Is always true when no startupProbe is defined.
     */
    started?: boolean;
    /**
     * Details about the container's current condition.
     */
    state?: ContainerState;
}
/**
 * DaemonEndpoint contains information about a single Daemon endpoint.
 */
export type DaemonEndpoint = {
    /**
     * Port number of the given endpoint.
     */
    Port: number;
}
/**
 * Represents downward API info for projecting into a projected volume. Note that this is identical to a downwardAPI volume source without the default mode.
 */
export type DownwardAPIProjection = {
    /**
     * Items is a list of DownwardAPIVolume file
     */
    items?: Array<DownwardAPIVolumeFile>;
}
/**
 * DownwardAPIVolumeFile represents information to create the file containing the pod field
 */
export type DownwardAPIVolumeFile = {
    /**
     * Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.
     */
    fieldRef?: ObjectFieldSelector;
    /**
     * Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
     */
    mode?: number;
    /**
     * Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'
     */
    path: string;
    /**
     * Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.
     */
    resourceFieldRef?: ResourceFieldSelector;
}
/**
 * DownwardAPIVolumeSource represents a volume containing downward API info. Downward API volumes support ownership management and SELinux relabeling.
 */
export type DownwardAPIVolumeSource = {
    /**
     * Optional: mode bits to use on created files by default. Must be a Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
     */
    defaultMode?: number;
    /**
     * Items is a list of downward API volume file
     */
    items?: Array<DownwardAPIVolumeFile>;
}
/**
 * Represents an empty directory for a pod. Empty directory volumes support ownership management and SELinux relabeling.
 */
export type EmptyDirVolumeSource = {
    /**
     * medium represents what type of storage medium should back this directory. The default is "" which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir
     */
    medium?: string;
    /**
     * sizeLimit is the total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: http://kubernetes.io/docs/user-guide/volumes#emptydir
     */
    sizeLimit?: Quantity;
}
/**
 * EndpointAddress is a tuple that describes single IP address.
 */
export type EndpointAddress = {
    /**
     * The Hostname of this endpoint
     */
    hostname?: string;
    /**
     * The IP of this endpoint. May not be loopback (127.0.0.0/8), link-local (169.254.0.0/16), or link-local multicast ((224.0.0.0/24). IPv6 is also accepted but not fully supported on all platforms. Also, certain kubernetes components, like kube-proxy, are not IPv6 ready.
     */
    ip: string;
    /**
     * Optional: Node hosting this endpoint. This can be used to determine endpoints local to a node.
     */
    nodeName?: string;
    /**
     * Reference to object providing the endpoint.
     */
    targetRef?: ObjectReference;
}
/**
 * EndpointPort is a tuple that describes a single port.
 */
export type EndpointPort = {
    /**
     * The application protocol for this port. This field follows standard Kubernetes label syntax. Un-prefixed names are reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names). Non-standard protocols should use prefixed names such as mycompany.com/my-custom-protocol.
     */
    appProtocol?: string;
    /**
     * The name of this port.  This must match the 'name' field in the corresponding ServicePort. Must be a DNS_LABEL. Optional only if one port is defined.
     */
    name?: string;
    /**
     * The port number of the endpoint.
     */
    port: number;
    /**
     * The IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.
     *
     *
     */
    protocol?: string;
}
/**
 * EndpointSubset is a group of addresses with a common set of ports. The expanded set of endpoints is the Cartesian product of Addresses x Ports. For example, given:
 *   {
 *     Addresses: [{"ip": "10.10.1.1"}, {"ip": "10.10.2.2"}],
 *     Ports:     [{"name": "a", "port": 8675}, {"name": "b", "port": 309}]
 *   }
 * The resulting set of endpoints can be viewed as:
 *     a: [ 10.10.1.1:8675, 10.10.2.2:8675 ],
 *     b: [ 10.10.1.1:309, 10.10.2.2:309 ]
 */
export type EndpointSubset = {
    /**
     * IP addresses which offer the related ports that are marked as ready. These endpoints should be considered safe for load balancers and clients to utilize.
     */
    addresses?: Array<EndpointAddress>;
    /**
     * IP addresses which offer the related ports but are not currently marked as ready because they have not yet finished starting, have recently failed a readiness check, or have recently failed a liveness check.
     */
    notReadyAddresses?: Array<EndpointAddress>;
    /**
     * Port numbers available on the related IP addresses.
     */
    ports?: Array<EndpointPort>;
}
/**
 * Endpoints is a collection of endpoints that implement the actual service. Example:
 *   Name: "mysvc",
 *   Subsets: [
 *     {
 *       Addresses: [{"ip": "10.10.1.1"}, {"ip": "10.10.2.2"}],
 *       Ports: [{"name": "a", "port": 8675}, {"name": "b", "port": 309}]
 *     },
 *     {
 *       Addresses: [{"ip": "10.10.3.3"}],
 *       Ports: [{"name": "a", "port": 93}, {"name": "b", "port": 76}]
 *     },
 *  ]
 */
export type Endpoints = {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "v1";
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "Endpoints";
    /**
     * Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     */
    metadata?: ObjectMeta;
    /**
     * The set of all endpoints is the union of all subsets. Addresses are placed into subsets according to the IPs they share. A single address with multiple ports, some of which are ready and some of which are not (because they come from different containers) will result in the address being displayed in different subsets for the different ports. No address will appear in both Addresses and NotReadyAddresses in the same subset. Sets of addresses and ports that comprise a service.
     */
    subsets?: Array<EndpointSubset>;
}
/**
 * EndpointsList is a list of endpoints.
 */
export type EndpointsList = {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "v1";
    /**
     * List of endpoints.
     */
    items: Array<Endpoints>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "EndpointsList";
    /**
     * Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    metadata?: ListMeta;
}
/**
 * EnvFromSource represents the source of a set of ConfigMaps
 */
export type EnvFromSource = {
    /**
     * The ConfigMap to select from
     */
    configMapRef?: ConfigMapEnvSource;
    /**
     * An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.
     */
    prefix?: string;
    /**
     * The Secret to select from
     */
    secretRef?: SecretEnvSource;
}
/**
 * EnvVar represents an environment variable present in a Container.
 */
export type EnvVar = {
    /**
     * Name of the environment variable. Must be a C_IDENTIFIER.
     */
    name: string;
    /**
     * Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)" will produce the string literal "$(VAR_NAME)". Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to "".
     */
    value?: string;
    /**
     * Source for the environment variable's value. Cannot be used if value is not empty.
     */
    valueFrom?: EnvVarSource;
}
/**
 * EnvVarSource represents a source for the value of an EnvVar.
 */
export type EnvVarSource = {
    /**
     * Selects a key of a ConfigMap.
     */
    configMapKeyRef?: ConfigMapKeySelector;
    /**
     * Selects a field of the pod: supports metadata.name, metadata.namespace, `metadata.labels['<KEY>']`, `metadata.annotations['<KEY>']`, spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.
     */
    fieldRef?: ObjectFieldSelector;
    /**
     * Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.
     */
    resourceFieldRef?: ResourceFieldSelector;
    /**
     * Selects a key of a secret in the pod's namespace
     */
    secretKeyRef?: SecretKeySelector;
}
/**
 * An EphemeralContainer is a temporary container that you may add to an existing Pod for user-initiated activities such as debugging. Ephemeral containers have no resource or scheduling guarantees, and they will not be restarted when they exit or when a Pod is removed or restarted. The kubelet may evict a Pod if an ephemeral container causes the Pod to exceed its resource allocation.
 *
 * To add an ephemeral container, use the ephemeralcontainers subresource of an existing Pod. Ephemeral containers may not be removed or restarted.
 *
 * This is a beta feature available on clusters that haven't disabled the EphemeralContainers feature gate.
 */
export type EphemeralContainer = {
    /**
     * Arguments to the entrypoint. The image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)" will produce the string literal "$(VAR_NAME)". Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
     */
    args?: Array<string>;
    /**
     * Entrypoint array. Not executed within a shell. The image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)" will produce the string literal "$(VAR_NAME)". Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
     */
    command?: Array<string>;
    /**
     * List of environment variables to set in the container. Cannot be updated.
     */
    env?: Array<EnvVar>;
    /**
     * List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated.
     */
    envFrom?: Array<EnvFromSource>;
    /**
     * Container image name. More info: https://kubernetes.io/docs/concepts/containers/images
     */
    image?: string;
    /**
     * Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images
     *
     *
     */
    imagePullPolicy?: string;
    /**
     * Lifecycle is not allowed for ephemeral containers.
     */
    lifecycle?: Lifecycle;
    /**
     * Probes are not allowed for ephemeral containers.
     */
    livenessProbe?: Probe;
    /**
     * Name of the ephemeral container specified as a DNS_LABEL. This name must be unique among all containers, init containers and ephemeral containers.
     */
    name: string;
    /**
     * Ports are not allowed for ephemeral containers.
     */
    ports?: Array<ContainerPort>;
    /**
     * Probes are not allowed for ephemeral containers.
     */
    readinessProbe?: Probe;
    /**
     * Resources are not allowed for ephemeral containers. Ephemeral containers use spare resources already allocated to the pod.
     */
    resources?: ResourceRequirements;
    /**
     * Optional: SecurityContext defines the security options the ephemeral container should be run with. If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext.
     */
    securityContext?: SecurityContext;
    /**
     * Probes are not allowed for ephemeral containers.
     */
    startupProbe?: Probe;
    /**
     * Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.
     */
    stdin?: boolean;
    /**
     * Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false
     */
    stdinOnce?: boolean;
    /**
     * If set, the name of the container from PodSpec that this ephemeral container targets. The ephemeral container will be run in the namespaces (IPC, PID, etc) of this container. If not set then the ephemeral container uses the namespaces configured in the Pod spec.
     *
     * The container runtime must implement support for this feature. If the runtime does not support namespace targeting then the result of setting this field is undefined.
     */
    targetContainerName?: string;
    /**
     * Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated.
     */
    terminationMessagePath?: string;
    /**
     * Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.
     *
     *
     */
    terminationMessagePolicy?: string;
    /**
     * Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false.
     */
    tty?: boolean;
    /**
     * volumeDevices is the list of block devices to be used by the container.
     */
    volumeDevices?: Array<VolumeDevice>;
    /**
     * Pod volumes to mount into the container's filesystem. Subpath mounts are not allowed for ephemeral containers. Cannot be updated.
     */
    volumeMounts?: Array<VolumeMount>;
    /**
     * Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.
     */
    workingDir?: string;
}
/**
 * Represents an ephemeral volume that is handled by a normal storage driver.
 */
export type EphemeralVolumeSource = {
    /**
     * Will be used to create a stand-alone PVC to provision the volume. The pod in which this EphemeralVolumeSource is embedded will be the owner of the PVC, i.e. the PVC will be deleted together with the pod.  The name of the PVC will be `<pod name>-<volume name>` where `<volume name>` is the name from the `PodSpec.Volumes` array entry. Pod validation will reject the pod if the concatenated name is not valid for a PVC (for example, too long).
     *
     * An existing PVC with that name that is not owned by the pod will *not* be used for the pod to avoid using an unrelated volume by mistake. Starting the pod is then blocked until the unrelated PVC is removed. If such a pre-created PVC is meant to be used by the pod, the PVC has to updated with an owner reference to the pod once the pod exists. Normally this should not be necessary, but it may be useful when manually reconstructing a broken cluster.
     *
     * This field is read-only and no changes will be made by Kubernetes to the PVC after it has been created.
     *
     * Required, must not be nil.
     */
    volumeClaimTemplate?: PersistentVolumeClaimTemplate;
}
/**
 * Event is a report of an event somewhere in the cluster.  Events have a limited retention time and triggers and messages may evolve with time.  Event consumers should not rely on the timing of an event with a given Reason reflecting a consistent underlying trigger, or the continued existence of events with that Reason.  Events should be treated as informative, best-effort, supplemental data.
 */
export type Event = {
    /**
     * What action was taken/failed regarding to the Regarding object.
     */
    action?: string;
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "v1";
    /**
     * The number of times this event has occurred.
     */
    count?: number;
    /**
     * Time when this Event was first observed.
     */
    eventTime?: MicroTime;
    /**
     * The time at which the event was first recorded. (Time of server receipt is in TypeMeta.)
     */
    firstTimestamp?: Time;
    /**
     * The object that this event is about.
     */
    involvedObject: ObjectReference;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "Event";
    /**
     * The time at which the most recent occurrence of this event was recorded.
     */
    lastTimestamp?: Time;
    /**
     * A human-readable description of the status of this operation.
     */
    message?: string;
    /**
     * Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     */
    metadata: ObjectMeta;
    /**
     * This should be a short, machine understandable string that gives the reason for the transition into the object's current status.
     */
    reason?: string;
    /**
     * Optional secondary object for more complex actions.
     */
    related?: ObjectReference;
    /**
     * Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
     */
    reportingComponent?: string;
    /**
     * ID of the controller instance, e.g. `kubelet-xyzf`.
     */
    reportingInstance?: string;
    /**
     * Data about the Event series this event represents or nil if it's a singleton Event.
     */
    series?: EventSeries;
    /**
     * The component reporting this event. Should be a short machine understandable string.
     */
    source?: EventSource;
    /**
     * Type of this event (Normal, Warning), new types could be added in the future
     */
    type?: string;
}
/**
 * EventList is a list of events.
 */
export type EventList = {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "v1";
    /**
     * List of events
     */
    items: Array<Event>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "EventList";
    /**
     * Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    metadata?: ListMeta;
}
/**
 * EventSeries contain information on series of events, i.e. thing that was/is happening continuously for some time.
 */
export type EventSeries = {
    /**
     * Number of occurrences in this series up to the last heartbeat time
     */
    count?: number;
    /**
     * Time of the last occurrence observed
     */
    lastObservedTime?: MicroTime;
}
/**
 * EventSource contains information for an event.
 */
export type EventSource = {
    /**
     * Component from which the event is generated.
     */
    component?: string;
    /**
     * Node name on which the event is generated.
     */
    host?: string;
}
/**
 * ExecAction describes a "run in container" action.
 */
export type ExecAction = {
    /**
     * Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.
     */
    command?: Array<string>;
}
/**
 * Represents a Fibre Channel volume. Fibre Channel volumes can only be mounted as read/write once. Fibre Channel volumes support ownership management and SELinux relabeling.
 */
export type FCVolumeSource = {
    /**
     * fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
     */
    fsType?: string;
    /**
     * lun is Optional: FC target lun number
     */
    lun?: number;
    /**
     * readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
     */
    readOnly?: boolean;
    /**
     * targetWWNs is Optional: FC target worldwide names (WWNs)
     */
    targetWWNs?: Array<string>;
    /**
     * wwids Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.
     */
    wwids?: Array<string>;
}
/**
 * FlexPersistentVolumeSource represents a generic persistent volume resource that is provisioned/attached using an exec based plugin.
 */
export type FlexPersistentVolumeSource = {
    /**
     * driver is the name of the driver to use for this volume.
     */
    driver: string;
    /**
     * fsType is the Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". The default filesystem depends on FlexVolume script.
     */
    fsType?: string;
    /**
     * options is Optional: this field holds extra command options if any.
     */
    options?: {
        [name: string]: string;
    };
    /**
     * readOnly is Optional: defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
     */
    readOnly?: boolean;
    /**
     * secretRef is Optional: SecretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.
     */
    secretRef?: SecretReference;
}
/**
 * FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.
 */
export type FlexVolumeSource = {
    /**
     * driver is the name of the driver to use for this volume.
     */
    driver: string;
    /**
     * fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". The default filesystem depends on FlexVolume script.
     */
    fsType?: string;
    /**
     * options is Optional: this field holds extra command options if any.
     */
    options?: {
        [name: string]: string;
    };
    /**
     * readOnly is Optional: defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
     */
    readOnly?: boolean;
    /**
     * secretRef is Optional: secretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.
     */
    secretRef?: LocalObjectReference;
}
/**
 * Represents a Flocker volume mounted by the Flocker agent. One and only one of datasetName and datasetUUID should be set. Flocker volumes do not support ownership management or SELinux relabeling.
 */
export type FlockerVolumeSource = {
    /**
     * datasetName is Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated
     */
    datasetName?: string;
    /**
     * datasetUUID is the UUID of the dataset. This is unique identifier of a Flocker dataset
     */
    datasetUUID?: string;
}
/**
 * Represents a Persistent Disk resource in Google Compute Engine.
 *
 * A GCE PD must exist before mounting to a container. The disk must also be in the same GCE project and zone as the kubelet. A GCE PD can only be mounted as read/write once or read-only many times. GCE PDs support ownership management and SELinux relabeling.
 */
export type GCEPersistentDiskVolumeSource = {
    /**
     * fsType is filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
     */
    fsType?: string;
    /**
     * partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as "1". Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
     */
    partition?: number;
    /**
     * pdName is unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
     */
    pdName: string;
    /**
     * readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
     */
    readOnly?: boolean;
}
export type GRPCAction = {
    /**
     * Port number of the gRPC service. Number must be in the range 1 to 65535.
     */
    port: number;
    /**
     * Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).
     *
     * If this is not specified, the default behavior is defined by gRPC.
     */
    service?: string;
}
/**
 * Represents a volume that is populated with the contents of a git repository. Git repo volumes do not support ownership management. Git repo volumes support SELinux relabeling.
 *
 * DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.
 */
export type GitRepoVolumeSource = {
    /**
     * directory is the target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.
     */
    directory?: string;
    /**
     * repository is the URL
     */
    repository: string;
    /**
     * revision is the commit hash for the specified revision.
     */
    revision?: string;
}
/**
 * Represents a Glusterfs mount that lasts the lifetime of a pod. Glusterfs volumes do not support ownership management or SELinux relabeling.
 */
export type GlusterfsPersistentVolumeSource = {
    /**
     * endpoints is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
     */
    endpoints: string;
    /**
     * endpointsNamespace is the namespace that contains Glusterfs endpoint. If this field is empty, the EndpointNamespace defaults to the same namespace as the bound PVC. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
     */
    endpointsNamespace?: string;
    /**
     * path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
     */
    path: string;
    /**
     * readOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
     */
    readOnly?: boolean;
}
/**
 * Represents a Glusterfs mount that lasts the lifetime of a pod. Glusterfs volumes do not support ownership management or SELinux relabeling.
 */
export type GlusterfsVolumeSource = {
    /**
     * endpoints is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
     */
    endpoints: string;
    /**
     * path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
     */
    path: string;
    /**
     * readOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
     */
    readOnly?: boolean;
}
/**
 * HTTPGetAction describes an action based on HTTP Get requests.
 */
export type HTTPGetAction = {
    /**
     * Host name to connect to, defaults to the pod IP. You probably want to set "Host" in httpHeaders instead.
     */
    host?: string;
    /**
     * Custom headers to set in the request. HTTP allows repeated headers.
     */
    httpHeaders?: Array<HTTPHeader>;
    /**
     * Path to access on the HTTP server.
     */
    path?: string;
    /**
     * Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
     */
    port: number | string;
    /**
     * Scheme to use for connecting to the host. Defaults to HTTP.
     *
     *
     */
    scheme?: string;
}
/**
 * HTTPHeader describes a custom header to be used in HTTP probes
 */
export type HTTPHeader = {
    /**
     * The header field name
     */
    name: string;
    /**
     * The header field value
     */
    value: string;
}
/**
 * HostAlias holds the mapping between IP and hostnames that will be injected as an entry in the pod's hosts file.
 */
export type HostAlias = {
    /**
     * Hostnames for the above IP address.
     */
    hostnames?: Array<string>;
    /**
     * IP address of the host file entry.
     */
    ip?: string;
}
/**
 * Represents a host path mapped into a pod. Host path volumes do not support ownership management or SELinux relabeling.
 */
export type HostPathVolumeSource = {
    /**
     * path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
     */
    path: string;
    /**
     * type for HostPath Volume Defaults to "" More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
     */
    type?: string;
}
/**
 * ISCSIPersistentVolumeSource represents an ISCSI disk. ISCSI volumes can only be mounted as read/write once. ISCSI volumes support ownership management and SELinux relabeling.
 */
export type ISCSIPersistentVolumeSource = {
    /**
     * chapAuthDiscovery defines whether support iSCSI Discovery CHAP authentication
     */
    chapAuthDiscovery?: boolean;
    /**
     * chapAuthSession defines whether support iSCSI Session CHAP authentication
     */
    chapAuthSession?: boolean;
    /**
     * fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi
     */
    fsType?: string;
    /**
     * initiatorName is the custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.
     */
    initiatorName?: string;
    /**
     * iqn is Target iSCSI Qualified Name.
     */
    iqn: string;
    /**
     * iscsiInterface is the interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).
     */
    iscsiInterface?: string;
    /**
     * lun is iSCSI Target Lun number.
     */
    lun: number;
    /**
     * portals is the iSCSI Target Portal List. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).
     */
    portals?: Array<string>;
    /**
     * readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.
     */
    readOnly?: boolean;
    /**
     * secretRef is the CHAP Secret for iSCSI target and initiator authentication
     */
    secretRef?: SecretReference;
    /**
     * targetPortal is iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).
     */
    targetPortal: string;
}
/**
 * Represents an ISCSI disk. ISCSI volumes can only be mounted as read/write once. ISCSI volumes support ownership management and SELinux relabeling.
 */
export type ISCSIVolumeSource = {
    /**
     * chapAuthDiscovery defines whether support iSCSI Discovery CHAP authentication
     */
    chapAuthDiscovery?: boolean;
    /**
     * chapAuthSession defines whether support iSCSI Session CHAP authentication
     */
    chapAuthSession?: boolean;
    /**
     * fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi
     */
    fsType?: string;
    /**
     * initiatorName is the custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.
     */
    initiatorName?: string;
    /**
     * iqn is the target iSCSI Qualified Name.
     */
    iqn: string;
    /**
     * iscsiInterface is the interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).
     */
    iscsiInterface?: string;
    /**
     * lun represents iSCSI Target Lun number.
     */
    lun: number;
    /**
     * portals is the iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).
     */
    portals?: Array<string>;
    /**
     * readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.
     */
    readOnly?: boolean;
    /**
     * secretRef is the CHAP Secret for iSCSI target and initiator authentication
     */
    secretRef?: LocalObjectReference;
    /**
     * targetPortal is iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).
     */
    targetPortal: string;
}
/**
 * Maps a string key to a path within a volume.
 */
export type KeyToPath = {
    /**
     * key is the key to project.
     */
    key: string;
    /**
     * mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
     */
    mode?: number;
    /**
     * path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.
     */
    path: string;
}
/**
 * Lifecycle describes actions that the management system should take in response to container lifecycle events. For the PostStart and PreStop lifecycle handlers, management of the container blocks until the action is complete, unless the container process fails, in which case the handler is aborted.
 */
export type Lifecycle = {
    /**
     * PostStart is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy. Other management of the container blocks until the hook completes. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks
     */
    postStart?: LifecycleHandler;
    /**
     * PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc. The handler is not called if the container crashes or exits. The Pod's termination grace period countdown begins before the PreStop hook is executed. Regardless of the outcome of the handler, the container will eventually terminate within the Pod's termination grace period (unless delayed by finalizers). Other management of the container blocks until the hook completes or until the termination grace period is reached. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks
     */
    preStop?: LifecycleHandler;
}
/**
 * LifecycleHandler defines a specific action that should be taken in a lifecycle hook. One and only one of the fields, except TCPSocket must be specified.
 */
export type LifecycleHandler = {
    /**
     * Exec specifies the action to take.
     */
    exec?: ExecAction;
    /**
     * HTTPGet specifies the http request to perform.
     */
    httpGet?: HTTPGetAction;
    /**
     * Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.
     */
    tcpSocket?: TCPSocketAction;
}
/**
 * LimitRange sets resource usage limits for each kind of resource in a Namespace.
 */
export type LimitRange = {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "v1";
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "LimitRange";
    /**
     * Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     */
    metadata?: ObjectMeta;
    /**
     * Spec defines the limits enforced. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     */
    spec?: LimitRangeSpec;
}
/**
 * LimitRangeItem defines a min/max usage limit for any resource that matches on kind.
 */
export type LimitRangeItem = {
    /**
     * Default resource requirement limit value by resource name if resource limit is omitted.
     */
    default?: {
        [name: string]: Quantity;
    };
    /**
     * DefaultRequest is the default resource requirement request value by resource name if resource request is omitted.
     */
    defaultRequest?: {
        [name: string]: Quantity;
    };
    /**
     * Max usage constraints on this kind by resource name.
     */
    max?: {
        [name: string]: Quantity;
    };
    /**
     * MaxLimitRequestRatio if specified, the named resource must have a request and limit that are both non-zero where limit divided by request is less than or equal to the enumerated value; this represents the max burst for the named resource.
     */
    maxLimitRequestRatio?: {
        [name: string]: Quantity;
    };
    /**
     * Min usage constraints on this kind by resource name.
     */
    min?: {
        [name: string]: Quantity;
    };
    /**
     * Type of resource that this limit applies to.
     */
    type: string;
}
/**
 * LimitRangeList is a list of LimitRange items.
 */
export type LimitRangeList = {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "v1";
    /**
     * Items is a list of LimitRange objects. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
     */
    items: Array<LimitRange>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "LimitRangeList";
    /**
     * Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    metadata?: ListMeta;
}
/**
 * LimitRangeSpec defines a min/max usage limit for resources that match on kind.
 */
export type LimitRangeSpec = {
    /**
     * Limits is the list of LimitRangeItem objects that are enforced.
     */
    limits: Array<LimitRangeItem>;
}
/**
 * LoadBalancerIngress represents the status of a load-balancer ingress point: traffic intended for the service should be sent to an ingress point.
 */
export type LoadBalancerIngress = {
    /**
     * Hostname is set for load-balancer ingress points that are DNS based (typically AWS load-balancers)
     */
    hostname?: string;
    /**
     * IP is set for load-balancer ingress points that are IP based (typically GCE or OpenStack load-balancers)
     */
    ip?: string;
    /**
     * Ports is a list of records of service ports If used, every port defined in the service should have an entry in it
     */
    ports?: Array<PortStatus>;
}
/**
 * LoadBalancerStatus represents the status of a load-balancer.
 */
export type LoadBalancerStatus = {
    /**
     * Ingress is a list containing ingress points for the load-balancer. Traffic intended for the service should be sent to these ingress points.
     */
    ingress?: Array<LoadBalancerIngress>;
}
/**
 * LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
 */
export type LocalObjectReference = {
    /**
     * Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
     */
    name?: string;
}
/**
 * Local represents directly-attached storage with node affinity (Beta feature)
 */
export type LocalVolumeSource = {
    /**
     * fsType is the filesystem type to mount. It applies only when the Path is a block device. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". The default value is to auto-select a filesystem if unspecified.
     */
    fsType?: string;
    /**
     * path of the full path to the volume on the node. It can be either a directory or block device (disk, partition, ...).
     */
    path: string;
}
/**
 * Represents an NFS mount that lasts the lifetime of a pod. NFS volumes do not support ownership management or SELinux relabeling.
 */
export type NFSVolumeSource = {
    /**
     * path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
     */
    path: string;
    /**
     * readOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
     */
    readOnly?: boolean;
    /**
     * server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
     */
    server: string;
}
/**
 * Namespace provides a scope for Names. Use of multiple namespaces is optional.
 */
export type Namespace = {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "v1";
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "Namespace";
    /**
     * Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     */
    metadata?: ObjectMeta;
    /**
     * Spec defines the behavior of the Namespace. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     */
    spec?: NamespaceSpec;
    /**
     * Status describes the current status of a Namespace. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     */
    status?: NamespaceStatus;
}
/**
 * NamespaceCondition contains details about state of namespace.
 */
export type NamespaceCondition = {
    lastTransitionTime?: Time;
    message?: string;
    reason?: string;
    /**
     * Status of the condition, one of True, False, Unknown.
     */
    status: string;
    /**
     * Type of namespace controller condition.
     */
    type: string;
}
/**
 * NamespaceList is a list of Namespaces.
 */
export type NamespaceList = {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "v1";
    /**
     * Items is the list of Namespace objects in the list. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
     */
    items: Array<Namespace>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "NamespaceList";
    /**
     * Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    metadata?: ListMeta;
}
/**
 * NamespaceSpec describes the attributes on a Namespace.
 */
export type NamespaceSpec = {
    /**
     * Finalizers is an opaque list of values that must be empty to permanently remove object from storage. More info: https://kubernetes.io/docs/tasks/administer-cluster/namespaces/
     */
    finalizers?: Array<string>;
}
/**
 * NamespaceStatus is information about the current status of a Namespace.
 */
export type NamespaceStatus = {
    /**
     * Represents the latest available observations of a namespace's current state.
     */
    conditions?: Array<NamespaceCondition>;
    /**
     * Phase is the current lifecycle phase of the namespace. More info: https://kubernetes.io/docs/tasks/administer-cluster/namespaces/
     *
     *
     */
    phase?: string;
}
/**
 * Node is a worker node in Kubernetes. Each node will have a unique identifier in the cache (i.e. in etcd).
 */
export type Node = {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "v1";
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "Node";
    /**
     * Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     */
    metadata?: ObjectMeta;
    /**
     * Spec defines the behavior of a node. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     */
    spec?: NodeSpec;
    /**
     * Most recently observed status of the node. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     */
    readonly status?: NodeStatus;
}
/**
 * NodeAddress contains information for the node's address.
 */
export type NodeAddress = {
    /**
     * The node address.
     */
    address: string;
    /**
     * Node address type, one of Hostname, ExternalIP or InternalIP.
     */
    type: string;
}
/**
 * Node affinity is a group of node affinity scheduling rules.
 */
export type NodeAffinity = {
    /**
     * The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding "weight" to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.
     */
    preferredDuringSchedulingIgnoredDuringExecution?: Array<PreferredSchedulingTerm>;
    /**
     * If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.
     */
    requiredDuringSchedulingIgnoredDuringExecution?: NodeSelector;
}
/**
 * NodeCondition contains condition information for a node.
 */
export type NodeCondition = {
    /**
     * Last time we got an update on a given condition.
     */
    lastHeartbeatTime?: Time;
    /**
     * Last time the condition transit from one status to another.
     */
    lastTransitionTime?: Time;
    /**
     * Human readable message indicating details about last transition.
     */
    message?: string;
    /**
     * (brief) reason for the condition's last transition.
     */
    reason?: string;
    /**
     * Status of the condition, one of True, False, Unknown.
     */
    status: string;
    /**
     * Type of node condition.
     */
    type: string;
}
/**
 * NodeConfigSource specifies a source of node configuration. Exactly one subfield (excluding metadata) must be non-nil. This API is deprecated since 1.22
 */
export type NodeConfigSource = {
    /**
     * ConfigMap is a reference to a Node's ConfigMap
     */
    configMap?: ConfigMapNodeConfigSource;
}
/**
 * NodeConfigStatus describes the status of the config assigned by Node.Spec.ConfigSource.
 */
export type NodeConfigStatus = {
    /**
     * Active reports the checkpointed config the node is actively using. Active will represent either the current version of the Assigned config, or the current LastKnownGood config, depending on whether attempting to use the Assigned config results in an error.
     */
    active?: NodeConfigSource;
    /**
     * Assigned reports the checkpointed config the node will try to use. When Node.Spec.ConfigSource is updated, the node checkpoints the associated config payload to local disk, along with a record indicating intended config. The node refers to this record to choose its config checkpoint, and reports this record in Assigned. Assigned only updates in the status after the record has been checkpointed to disk. When the Kubelet is restarted, it tries to make the Assigned config the Active config by loading and validating the checkpointed payload identified by Assigned.
     */
    assigned?: NodeConfigSource;
    /**
     * Error describes any problems reconciling the Spec.ConfigSource to the Active config. Errors may occur, for example, attempting to checkpoint Spec.ConfigSource to the local Assigned record, attempting to checkpoint the payload associated with Spec.ConfigSource, attempting to load or validate the Assigned config, etc. Errors may occur at different points while syncing config. Earlier errors (e.g. download or checkpointing errors) will not result in a rollback to LastKnownGood, and may resolve across Kubelet retries. Later errors (e.g. loading or validating a checkpointed config) will result in a rollback to LastKnownGood. In the latter case, it is usually possible to resolve the error by fixing the config assigned in Spec.ConfigSource. You can find additional information for debugging by searching the error message in the Kubelet log. Error is a human-readable description of the error state; machines can check whether or not Error is empty, but should not rely on the stability of the Error text across Kubelet versions.
     */
    error?: string;
    /**
     * LastKnownGood reports the checkpointed config the node will fall back to when it encounters an error attempting to use the Assigned config. The Assigned config becomes the LastKnownGood config when the node determines that the Assigned config is stable and correct. This is currently implemented as a 10-minute soak period starting when the local record of Assigned config is updated. If the Assigned config is Active at the end of this period, it becomes the LastKnownGood. Note that if Spec.ConfigSource is reset to nil (use local defaults), the LastKnownGood is also immediately reset to nil, because the local default config is always assumed good. You should not make assumptions about the node's method of determining config stability and correctness, as this may change or become configurable in the future.
     */
    lastKnownGood?: NodeConfigSource;
}
/**
 * NodeDaemonEndpoints lists ports opened by daemons running on the Node.
 */
export type NodeDaemonEndpoints = {
    /**
     * Endpoint on which Kubelet is listening.
     */
    kubeletEndpoint?: DaemonEndpoint;
}
/**
 * NodeList is the whole list of all Nodes which have been registered with master.
 */
export type NodeList = {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "v1";
    /**
     * List of nodes
     */
    items: Array<Node>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "NodeList";
    /**
     * Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    metadata?: ListMeta;
}
/**
 * A node selector represents the union of the results of one or more label queries over a set of nodes; that is, it represents the OR of the selectors represented by the node selector terms.
 */
export type NodeSelector = {
    /**
     * Required. A list of node selector terms. The terms are ORed.
     */
    nodeSelectorTerms: Array<NodeSelectorTerm>;
}
/**
 * A node selector requirement is a selector that contains values, a key, and an operator that relates the key and values.
 */
export type NodeSelectorRequirement = {
    /**
     * The label key that the selector applies to.
     */
    key: string;
    /**
     * Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.
     *
     *
     */
    operator: string;
    /**
     * An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.
     */
    values?: Array<string>;
}
/**
 * A null or empty node selector term matches no objects. The requirements of them are ANDed. The TopologySelectorTerm type implements a subset of the NodeSelectorTerm.
 */
export type NodeSelectorTerm = {
    /**
     * A list of node selector requirements by node's labels.
     */
    matchExpressions?: Array<NodeSelectorRequirement>;
    /**
     * A list of node selector requirements by node's fields.
     */
    matchFields?: Array<NodeSelectorRequirement>;
}
/**
 * NodeSpec describes the attributes that a node is created with.
 */
export type NodeSpec = {
    /**
     * Deprecated: Previously used to specify the source of the node's configuration for the DynamicKubeletConfig feature. This feature is removed from Kubelets as of 1.24 and will be fully removed in 1.26.
     */
    configSource?: NodeConfigSource;
    /**
     * Deprecated. Not all kubelets will set this field. Remove field after 1.13. see: https://issues.k8s.io/61966
     */
    externalID?: string;
    /**
     * PodCIDR represents the pod IP range assigned to the node.
     */
    podCIDR?: string;
    /**
     * podCIDRs represents the IP ranges assigned to the node for usage by Pods on that node. If this field is specified, the 0th entry must match the podCIDR field. It may contain at most 1 value for each of IPv4 and IPv6.
     */
    podCIDRs?: Array<string>;
    /**
     * ID of the node assigned by the cloud provider in the format: <ProviderName>://<ProviderSpecificNodeID>
     */
    providerID?: string;
    /**
     * If specified, the node's taints.
     */
    taints?: Array<Taint>;
    /**
     * Unschedulable controls node schedulability of new pods. By default, node is schedulable. More info: https://kubernetes.io/docs/concepts/nodes/node/#manual-node-administration
     */
    unschedulable?: boolean;
}
/**
 * NodeStatus is information about the current status of a node.
 */
export type NodeStatus = {
    /**
     * List of addresses reachable to the node. Queried from cloud provider, if available. More info: https://kubernetes.io/docs/concepts/nodes/node/#addresses Note: This field is declared as mergeable, but the merge key is not sufficiently unique, which can cause data corruption when it is merged. Callers should instead use a full-replacement patch. See http://pr.k8s.io/79391 for an example.
     */
    addresses?: Array<NodeAddress>;
    /**
     * Allocatable represents the resources of a node that are available for scheduling. Defaults to Capacity.
     */
    allocatable?: {
        [name: string]: Quantity;
    };
    /**
     * Capacity represents the total resources of a node. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity
     */
    capacity?: {
        [name: string]: Quantity;
    };
    /**
     * Conditions is an array of current observed node conditions. More info: https://kubernetes.io/docs/concepts/nodes/node/#condition
     */
    conditions?: Array<NodeCondition>;
    /**
     * Status of the config assigned to the node via the dynamic Kubelet config feature.
     */
    config?: NodeConfigStatus;
    /**
     * Endpoints of daemons running on the Node.
     */
    daemonEndpoints?: NodeDaemonEndpoints;
    /**
     * List of container images on this node
     */
    images?: Array<ContainerImage>;
    /**
     * Set of ids/uuids to uniquely identify the node. More info: https://kubernetes.io/docs/concepts/nodes/node/#info
     */
    nodeInfo?: NodeSystemInfo;
    /**
     * NodePhase is the recently observed lifecycle phase of the node. More info: https://kubernetes.io/docs/concepts/nodes/node/#phase The field is never populated, and now is deprecated.
     *
     *
     */
    phase?: string;
    /**
     * List of volumes that are attached to the node.
     */
    volumesAttached?: Array<AttachedVolume>;
    /**
     * List of attachable volumes in use (mounted) by the node.
     */
    volumesInUse?: Array<string>;
}
/**
 * NodeSystemInfo is a set of ids/uuids to uniquely identify the node.
 */
export type NodeSystemInfo = {
    /**
     * The Architecture reported by the node
     */
    architecture: string;
    /**
     * Boot ID reported by the node.
     */
    bootID: string;
    /**
     * ContainerRuntime Version reported by the node through runtime remote API (e.g. containerd://1.4.2).
     */
    containerRuntimeVersion: string;
    /**
     * Kernel Version reported by the node from 'uname -r' (e.g. 3.16.0-0.bpo.4-amd64).
     */
    kernelVersion: string;
    /**
     * KubeProxy Version reported by the node.
     */
    kubeProxyVersion: string;
    /**
     * Kubelet Version reported by the node.
     */
    kubeletVersion: string;
    /**
     * MachineID reported by the node. For unique machine identification in the cluster this field is preferred. Learn more from man(5) machine-id: http://man7.org/linux/man-pages/man5/machine-id.5.html
     */
    machineID: string;
    /**
     * The Operating System reported by the node
     */
    operatingSystem: string;
    /**
     * OS Image reported by the node from /etc/os-release (e.g. Debian GNU/Linux 7 (wheezy)).
     */
    osImage: string;
    /**
     * SystemUUID reported by the node. For unique machine identification MachineID is preferred. This field is specific to Red Hat hosts https://access.redhat.com/documentation/en-us/red_hat_subscription_management/1/html/rhsm/uuid
     */
    systemUUID: string;
}
/**
 * ObjectFieldSelector selects an APIVersioned field of an object.
 */
export type ObjectFieldSelector = {
    /**
     * Version of the schema the FieldPath is written in terms of, defaults to "v1".
     */
    apiVersion?: string;
    /**
     * Path of the field to select in the specified API version.
     */
    fieldPath: string;
}
/**
 * ObjectReference contains enough information to let you inspect or modify the referred object.
 */
export type ObjectReference = {
    /**
     * API version of the referent.
     */
    apiVersion?: string;
    /**
     * If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: "spec.containers{name}" (where "name" refers to the name of the container that triggered the event) or if no container name is specified "spec.containers[2]" (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.
     */
    fieldPath?: string;
    /**
     * Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: string;
    /**
     * Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
     */
    name?: string;
    /**
     * Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
     */
    namespace?: string;
    /**
     * Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
     */
    resourceVersion?: string;
    /**
     * UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids
     */
    uid?: string;
}
/**
 * PersistentVolume (PV) is a storage resource provisioned by an administrator. It is analogous to a node. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes
 */
export type PersistentVolume = {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "v1";
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "PersistentVolume";
    /**
     * Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     */
    metadata?: ObjectMeta;
    /**
     * spec defines a specification of a persistent volume owned by the cluster. Provisioned by an administrator. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistent-volumes
     */
    spec?: PersistentVolumeSpec;
    /**
     * status represents the current information/status for the persistent volume. Populated by the system. Read-only. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistent-volumes
     */
    readonly status?: PersistentVolumeStatus;
}
/**
 * PersistentVolumeClaim is a user's request for and claim to a persistent volume
 */
export type PersistentVolumeClaim = {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "v1";
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "PersistentVolumeClaim";
    /**
     * Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     */
    metadata?: ObjectMeta;
    /**
     * spec defines the desired characteristics of a volume requested by a pod author. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
     */
    spec?: PersistentVolumeClaimSpec;
    /**
     * status represents the current information/status of a persistent volume claim. Read-only. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
     */
    readonly status?: PersistentVolumeClaimStatus;
}
/**
 * PersistentVolumeClaimCondition contails details about state of pvc
 */
export type PersistentVolumeClaimCondition = {
    /**
     * lastProbeTime is the time we probed the condition.
     */
    lastProbeTime?: Time;
    /**
     * lastTransitionTime is the time the condition transitioned from one status to another.
     */
    lastTransitionTime?: Time;
    /**
     * message is the human-readable message indicating details about last transition.
     */
    message?: string;
    /**
     * reason is a unique, this should be a short, machine understandable string that gives the reason for condition's last transition. If it reports "ResizeStarted" that means the underlying persistent volume is being resized.
     */
    reason?: string;
    status: string;
    type: string;
}
/**
 * PersistentVolumeClaimList is a list of PersistentVolumeClaim items.
 */
export type PersistentVolumeClaimList = {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "v1";
    /**
     * items is a list of persistent volume claims. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
     */
    items: Array<PersistentVolumeClaim>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "PersistentVolumeClaimList";
    /**
     * Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    metadata?: ListMeta;
}
/**
 * PersistentVolumeClaimSpec describes the common attributes of storage devices and allows a Source for provider-specific attributes
 */
export type PersistentVolumeClaimSpec = {
    /**
     * accessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1
     */
    accessModes?: Array<string>;
    /**
     * dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. If the AnyVolumeDataSource feature gate is enabled, this field will always have the same contents as the DataSourceRef field.
     */
    dataSource?: TypedLocalObjectReference;
    /**
     * dataSourceRef specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any local object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the DataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, both fields (DataSource and DataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. There are two important differences between DataSource and DataSourceRef: * While DataSource only allows two specific types of objects, DataSourceRef
     *   allows any non-core object, as well as PersistentVolumeClaim objects.
     * * While DataSource ignores disallowed values (dropping them), DataSourceRef
     *   preserves all values, and generates an error if a disallowed value is
     *   specified.
     * (Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.
     */
    dataSourceRef?: TypedLocalObjectReference;
    /**
     * resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources
     */
    resources?: ResourceRequirements;
    /**
     * selector is a label query over volumes to consider for binding.
     */
    selector?: LabelSelector;
    /**
     * storageClassName is the name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1
     */
    storageClassName?: string;
    /**
     * volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.
     */
    volumeMode?: string;
    /**
     * volumeName is the binding reference to the PersistentVolume backing this claim.
     */
    volumeName?: string;
}
/**
 * PersistentVolumeClaimStatus is the current status of a persistent volume claim.
 */
export type PersistentVolumeClaimStatus = {
    /**
     * accessModes contains the actual access modes the volume backing the PVC has. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1
     */
    accessModes?: Array<string>;
    /**
     * allocatedResources is the storage resource within AllocatedResources tracks the capacity allocated to a PVC. It may be larger than the actual capacity when a volume expansion operation is requested. For storage quota, the larger value from allocatedResources and PVC.spec.resources is used. If allocatedResources is not set, PVC.spec.resources alone is used for quota calculation. If a volume expansion capacity request is lowered, allocatedResources is only lowered if there are no expansion operations in progress and if the actual volume capacity is equal or lower than the requested capacity. This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.
     */
    allocatedResources?: {
        [name: string]: Quantity;
    };
    /**
     * capacity represents the actual resources of the underlying volume.
     */
    capacity?: {
        [name: string]: Quantity;
    };
    /**
     * conditions is the current Condition of persistent volume claim. If underlying persistent volume is being resized then the Condition will be set to 'ResizeStarted'.
     */
    conditions?: Array<PersistentVolumeClaimCondition>;
    /**
     * phase represents the current phase of PersistentVolumeClaim.
     *
     *
     */
    phase?: string;
    /**
     * resizeStatus stores status of resize operation. ResizeStatus is not set by default but when expansion is complete resizeStatus is set to empty string by resize controller or kubelet. This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.
     */
    resizeStatus?: string;
}
/**
 * PersistentVolumeClaimTemplate is used to produce PersistentVolumeClaim objects as part of an EphemeralVolumeSource.
 */
export type PersistentVolumeClaimTemplate = {
    /**
     * May contain labels and annotations that will be copied into the PVC when creating it. No other fields are allowed and will be rejected during validation.
     */
    metadata?: ObjectMeta;
    /**
     * The specification for the PersistentVolumeClaim. The entire content is copied unchanged into the PVC that gets created from this template. The same fields as in a PersistentVolumeClaim are also valid here.
     */
    spec: PersistentVolumeClaimSpec;
}
/**
 * PersistentVolumeClaimVolumeSource references the user's PVC in the same namespace. This volume finds the bound PV and mounts that volume for the pod. A PersistentVolumeClaimVolumeSource is, essentially, a wrapper around another type of volume that is owned by someone else (the system).
 */
export type PersistentVolumeClaimVolumeSource = {
    /**
     * claimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
     */
    claimName: string;
    /**
     * readOnly Will force the ReadOnly setting in VolumeMounts. Default false.
     */
    readOnly?: boolean;
}
/**
 * PersistentVolumeList is a list of PersistentVolume items.
 */
export type PersistentVolumeList = {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "v1";
    /**
     * items is a list of persistent volumes. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes
     */
    items: Array<PersistentVolume>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "PersistentVolumeList";
    /**
     * Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    metadata?: ListMeta;
}
/**
 * PersistentVolumeSpec is the specification of a persistent volume.
 */
export type PersistentVolumeSpec = {
    /**
     * accessModes contains all ways the volume can be mounted. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes
     */
    accessModes?: Array<string>;
    /**
     * awsElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
     */
    awsElasticBlockStore?: AWSElasticBlockStoreVolumeSource;
    /**
     * azureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
     */
    azureDisk?: AzureDiskVolumeSource;
    /**
     * azureFile represents an Azure File Service mount on the host and bind mount to the pod.
     */
    azureFile?: AzureFilePersistentVolumeSource;
    /**
     * capacity is the description of the persistent volume's resources and capacity. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity
     */
    capacity?: {
        [name: string]: Quantity;
    };
    /**
     * cephFS represents a Ceph FS mount on the host that shares a pod's lifetime
     */
    cephfs?: CephFSPersistentVolumeSource;
    /**
     * cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
     */
    cinder?: CinderPersistentVolumeSource;
    /**
     * claimRef is part of a bi-directional binding between PersistentVolume and PersistentVolumeClaim. Expected to be non-nil when bound. claim.VolumeName is the authoritative bind between PV and PVC. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#binding
     */
    claimRef?: ObjectReference;
    /**
     * csi represents storage that is handled by an external CSI driver (Beta feature).
     */
    csi?: CSIPersistentVolumeSource;
    /**
     * fc represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.
     */
    fc?: FCVolumeSource;
    /**
     * flexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.
     */
    flexVolume?: FlexPersistentVolumeSource;
    /**
     * flocker represents a Flocker volume attached to a kubelet's host machine and exposed to the pod for its usage. This depends on the Flocker control service being running
     */
    flocker?: FlockerVolumeSource;
    /**
     * gcePersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. Provisioned by an admin. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
     */
    gcePersistentDisk?: GCEPersistentDiskVolumeSource;
    /**
     * glusterfs represents a Glusterfs volume that is attached to a host and exposed to the pod. Provisioned by an admin. More info: https://examples.k8s.io/volumes/glusterfs/README.md
     */
    glusterfs?: GlusterfsPersistentVolumeSource;
    /**
     * hostPath represents a directory on the host. Provisioned by a developer or tester. This is useful for single-node development and testing only! On-host storage is not supported in any way and WILL NOT WORK in a multi-node cluster. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
     */
    hostPath?: HostPathVolumeSource;
    /**
     * iscsi represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. Provisioned by an admin.
     */
    iscsi?: ISCSIPersistentVolumeSource;
    /**
     * local represents directly-attached storage with node affinity
     */
    local?: LocalVolumeSource;
    /**
     * mountOptions is the list of mount options, e.g. ["ro", "soft"]. Not validated - mount will simply fail if one is invalid. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#mount-options
     */
    mountOptions?: Array<string>;
    /**
     * nfs represents an NFS mount on the host. Provisioned by an admin. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
     */
    nfs?: NFSVolumeSource;
    /**
     * nodeAffinity defines constraints that limit what nodes this volume can be accessed from. This field influences the scheduling of pods that use this volume.
     */
    nodeAffinity?: VolumeNodeAffinity;
    /**
     * persistentVolumeReclaimPolicy defines what happens to a persistent volume when released from its claim. Valid options are Retain (default for manually created PersistentVolumes), Delete (default for dynamically provisioned PersistentVolumes), and Recycle (deprecated). Recycle must be supported by the volume plugin underlying this PersistentVolume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#reclaiming
     *
     *
     */
    persistentVolumeReclaimPolicy?: string;
    /**
     * photonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine
     */
    photonPersistentDisk?: PhotonPersistentDiskVolumeSource;
    /**
     * portworxVolume represents a portworx volume attached and mounted on kubelets host machine
     */
    portworxVolume?: PortworxVolumeSource;
    /**
     * quobyte represents a Quobyte mount on the host that shares a pod's lifetime
     */
    quobyte?: QuobyteVolumeSource;
    /**
     * rbd represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md
     */
    rbd?: RBDPersistentVolumeSource;
    /**
     * scaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.
     */
    scaleIO?: ScaleIOPersistentVolumeSource;
    /**
     * storageClassName is the name of StorageClass to which this persistent volume belongs. Empty value means that this volume does not belong to any StorageClass.
     */
    storageClassName?: string;
    /**
     * storageOS represents a StorageOS volume that is attached to the kubelet's host machine and mounted into the pod More info: https://examples.k8s.io/volumes/storageos/README.md
     */
    storageos?: StorageOSPersistentVolumeSource;
    /**
     * volumeMode defines if a volume is intended to be used with a formatted filesystem or to remain in raw block state. Value of Filesystem is implied when not included in spec.
     */
    volumeMode?: string;
    /**
     * vsphereVolume represents a vSphere volume attached and mounted on kubelets host machine
     */
    vsphereVolume?: VsphereVirtualDiskVolumeSource;
}
/**
 * PersistentVolumeStatus is the current status of a persistent volume.
 */
export type PersistentVolumeStatus = {
    /**
     * message is a human-readable message indicating details about why the volume is in this state.
     */
    message?: string;
    /**
     * phase indicates if a volume is available, bound to a claim, or released by a claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#phase
     *
     *
     */
    phase?: string;
    /**
     * reason is a brief CamelCase string that describes any failure and is meant for machine parsing and tidy display in the CLI.
     */
    reason?: string;
}
/**
 * Represents a Photon Controller persistent disk resource.
 */
export type PhotonPersistentDiskVolumeSource = {
    /**
     * fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
     */
    fsType?: string;
    /**
     * pdID is the ID that identifies Photon Controller persistent disk
     */
    pdID: string;
}
/**
 * Pod is a collection of containers that can run on a host. This resource is created by clients and scheduled onto hosts.
 */
export type Pod = {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "v1";
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "Pod";
    /**
     * Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     */
    metadata?: ObjectMeta;
    /**
     * Specification of the desired behavior of the pod. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     */
    spec?: PodSpec;
    /**
     * Most recently observed status of the pod. This data may not be up to date. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     */
    readonly status?: PodStatus;
}
/**
 * Pod affinity is a group of inter pod affinity scheduling rules.
 */
export type PodAffinity = {
    /**
     * The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding "weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.
     */
    preferredDuringSchedulingIgnoredDuringExecution?: Array<WeightedPodAffinityTerm>;
    /**
     * If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.
     */
    requiredDuringSchedulingIgnoredDuringExecution?: Array<PodAffinityTerm>;
}
/**
 * Defines a set of pods (namely those matching the labelSelector relative to the given namespace(s)) that this pod should be co-located (affinity) or not co-located (anti-affinity) with, where co-located is defined as running on a node whose value of the label with key <topologyKey> matches that of any node on which a pod of the set of pods is running
 */
export type PodAffinityTerm = {
    /**
     * A label query over a set of resources, in this case pods.
     */
    labelSelector?: LabelSelector;
    /**
     * A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means "this pod's namespace". An empty selector ({}) matches all namespaces.
     */
    namespaceSelector?: LabelSelector;
    /**
     * namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means "this pod's namespace".
     */
    namespaces?: Array<string>;
    /**
     * This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.
     */
    topologyKey: string;
}
/**
 * Pod anti affinity is a group of inter pod anti affinity scheduling rules.
 */
export type PodAntiAffinity = {
    /**
     * The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding "weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.
     */
    preferredDuringSchedulingIgnoredDuringExecution?: Array<WeightedPodAffinityTerm>;
    /**
     * If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.
     */
    requiredDuringSchedulingIgnoredDuringExecution?: Array<PodAffinityTerm>;
}
/**
 * PodCondition contains details for the current condition of this pod.
 */
export type PodCondition = {
    /**
     * Last time we probed the condition.
     */
    lastProbeTime?: Time;
    /**
     * Last time the condition transitioned from one status to another.
     */
    lastTransitionTime?: Time;
    /**
     * Human-readable message indicating details about last transition.
     */
    message?: string;
    /**
     * Unique, one-word, CamelCase reason for the condition's last transition.
     */
    reason?: string;
    /**
     * Status is the status of the condition. Can be True, False, Unknown. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions
     */
    status: string;
    /**
     * Type is the type of the condition. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions
     */
    type: string;
}
/**
 * PodDNSConfig defines the DNS parameters of a pod in addition to those generated from DNSPolicy.
 */
export type PodDNSConfig = {
    /**
     * A list of DNS name server IP addresses. This will be appended to the base nameservers generated from DNSPolicy. Duplicated nameservers will be removed.
     */
    nameservers?: Array<string>;
    /**
     * A list of DNS resolver options. This will be merged with the base options generated from DNSPolicy. Duplicated entries will be removed. Resolution options given in Options will override those that appear in the base DNSPolicy.
     */
    options?: Array<PodDNSConfigOption>;
    /**
     * A list of DNS search domains for host-name lookup. This will be appended to the base search paths generated from DNSPolicy. Duplicated search paths will be removed.
     */
    searches?: Array<string>;
}
/**
 * PodDNSConfigOption defines DNS resolver options of a pod.
 */
export type PodDNSConfigOption = {
    /**
     * Required.
     */
    name?: string;
    value?: string;
}
/**
 * IP address information for entries in the (plural) PodIPs field. Each entry includes:
 *    IP: An IP address allocated to the pod. Routable at least within the cluster.
 */
export type PodIP = {
    /**
     * ip is an IP address (IPv4 or IPv6) assigned to the pod
     */
    ip?: string;
}
/**
 * PodList is a list of Pods.
 */
export type PodList = {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "v1";
    /**
     * List of pods. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
     */
    items: Array<Pod>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "PodList";
    /**
     * Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    metadata?: ListMeta;
}
/**
 * PodOS defines the OS parameters of a pod.
 */
export type PodOS = {
    /**
     * Name is the name of the operating system. The currently supported values are linux and windows. Additional value may be defined in future and can be one of: https://github.com/opencontainers/runtime-spec/blob/master/config.md#platform-specific-configuration Clients should expect to handle additional values and treat unrecognized values in this field as os: null
     */
    name: string;
}
/**
 * PodReadinessGate contains the reference to a pod condition
 */
export type PodReadinessGate = {
    /**
     * ConditionType refers to a condition in the pod's condition list with matching type.
     */
    conditionType: string;
}
/**
 * PodSecurityContext holds pod-level security attributes and common container settings. Some fields are also present in container.securityContext.  Field values of container.securityContext take precedence over field values of PodSecurityContext.
 */
export type PodSecurityContext = {
    /**
     * A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod:
     *
     * 1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw----
     *
     * If unset, the Kubelet will not modify the ownership and permissions of any volume. Note that this field cannot be set when spec.os.name is windows.
     */
    fsGroup?: number;
    /**
     * fsGroupChangePolicy defines behavior of changing ownership and permission of the volume before being exposed inside Pod. This field will only apply to volume types which support fsGroup based ownership(and permissions). It will have no effect on ephemeral volume types such as: secret, configmaps and emptydir. Valid values are "OnRootMismatch" and "Always". If not specified, "Always" is used. Note that this field cannot be set when spec.os.name is windows.
     */
    fsGroupChangePolicy?: string;
    /**
     * The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.
     */
    runAsGroup?: number;
    /**
     * Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
     */
    runAsNonRoot?: boolean;
    /**
     * The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.
     */
    runAsUser?: number;
    /**
     * The SELinux context to be applied to all containers. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.
     */
    seLinuxOptions?: SELinuxOptions;
    /**
     * The seccomp options to use by the containers in this pod. Note that this field cannot be set when spec.os.name is windows.
     */
    seccompProfile?: SeccompProfile;
    /**
     * A list of groups applied to the first process run in each container, in addition to the container's primary GID.  If unspecified, no groups will be added to any container. Note that this field cannot be set when spec.os.name is windows.
     */
    supplementalGroups?: Array<number>;
    /**
     * Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupported sysctls (by the container runtime) might fail to launch. Note that this field cannot be set when spec.os.name is windows.
     */
    sysctls?: Array<Sysctl>;
    /**
     * The Windows specific settings applied to all containers. If unspecified, the options within a container's SecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.
     */
    windowsOptions?: WindowsSecurityContextOptions;
}
/**
 * PodSpec is a description of a pod.
 */
export type PodSpec = {
    /**
     * Optional duration in seconds the pod may be active on the node relative to StartTime before the system will actively try to mark it failed and kill associated containers. Value must be a positive integer.
     */
    activeDeadlineSeconds?: number;
    /**
     * If specified, the pod's scheduling constraints
     */
    affinity?: Affinity;
    /**
     * AutomountServiceAccountToken indicates whether a service account token should be automatically mounted.
     */
    automountServiceAccountToken?: boolean;
    /**
     * List of containers belonging to the pod. Containers cannot currently be added or removed. There must be at least one container in a Pod. Cannot be updated.
     */
    containers: Array<Container>;
    /**
     * Specifies the DNS parameters of a pod. Parameters specified here will be merged to the generated DNS configuration based on DNSPolicy.
     */
    dnsConfig?: PodDNSConfig;
    /**
     * Set DNS policy for the pod. Defaults to "ClusterFirst". Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'. DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy. To have DNS options set along with hostNetwork, you have to specify DNS policy explicitly to 'ClusterFirstWithHostNet'.
     *
     *
     */
    dnsPolicy?: string;
    /**
     * EnableServiceLinks indicates whether information about services should be injected into pod's environment variables, matching the syntax of Docker links. Optional: Defaults to true.
     */
    enableServiceLinks?: boolean;
    /**
     * List of ephemeral containers run in this pod. Ephemeral containers may be run in an existing pod to perform user-initiated actions such as debugging. This list cannot be specified when creating a pod, and it cannot be modified by updating the pod spec. In order to add an ephemeral container to an existing pod, use the pod's ephemeralcontainers subresource. This field is beta-level and available on clusters that haven't disabled the EphemeralContainers feature gate.
     */
    ephemeralContainers?: Array<EphemeralContainer>;
    /**
     * HostAliases is an optional list of hosts and IPs that will be injected into the pod's hosts file if specified. This is only valid for non-hostNetwork pods.
     */
    hostAliases?: Array<HostAlias>;
    /**
     * Use the host's ipc namespace. Optional: Default to false.
     */
    hostIPC?: boolean;
    /**
     * Host networking requested for this pod. Use the host's network namespace. If this option is set, the ports that will be used must be specified. Default to false.
     */
    hostNetwork?: boolean;
    /**
     * Use the host's pid namespace. Optional: Default to false.
     */
    hostPID?: boolean;
    /**
     * Specifies the hostname of the Pod If not specified, the pod's hostname will be set to a system-defined value.
     */
    hostname?: string;
    /**
     * ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec. If specified, these secrets will be passed to individual puller implementations for them to use. More info: https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod
     */
    imagePullSecrets?: Array<LocalObjectReference>;
    /**
     * List of initialization containers belonging to the pod. Init containers are executed in order prior to containers being started. If any init container fails, the pod is considered to have failed and is handled according to its restartPolicy. The name for an init container or normal container must be unique among all containers. Init containers may not have Lifecycle actions, Readiness probes, Liveness probes, or Startup probes. The resourceRequirements of an init container are taken into account during scheduling by finding the highest request/limit for each resource type, and then using the max of of that value or the sum of the normal containers. Limits are applied to init containers in a similar fashion. Init containers cannot currently be added or removed. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/
     */
    initContainers?: Array<Container>;
    /**
     * NodeName is a request to schedule this pod onto a specific node. If it is non-empty, the scheduler simply schedules this pod onto that node, assuming that it fits resource requirements.
     */
    nodeName?: string;
    /**
     * NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
     */
    nodeSelector?: {
        [name: string]: string;
    };
    /**
     * Specifies the OS of the containers in the pod. Some pod and container fields are restricted if this is set.
     *
     * If the OS field is set to linux, the following fields must be unset: -securityContext.windowsOptions
     *
     * If the OS field is set to windows, following fields must be unset: - spec.hostPID - spec.hostIPC - spec.securityContext.seLinuxOptions - spec.securityContext.seccompProfile - spec.securityContext.fsGroup - spec.securityContext.fsGroupChangePolicy - spec.securityContext.sysctls - spec.shareProcessNamespace - spec.securityContext.runAsUser - spec.securityContext.runAsGroup - spec.securityContext.supplementalGroups - spec.containers[*].securityContext.seLinuxOptions - spec.containers[*].securityContext.seccompProfile - spec.containers[*].securityContext.capabilities - spec.containers[*].securityContext.readOnlyRootFilesystem - spec.containers[*].securityContext.privileged - spec.containers[*].securityContext.allowPrivilegeEscalation - spec.containers[*].securityContext.procMount - spec.containers[*].securityContext.runAsUser - spec.containers[*].securityContext.runAsGroup This is a beta field and requires the IdentifyPodOS feature
     */
    os?: PodOS;
    /**
     * Overhead represents the resource overhead associated with running a pod for a given RuntimeClass. This field will be autopopulated at admission time by the RuntimeClass admission controller. If the RuntimeClass admission controller is enabled, overhead must not be set in Pod create requests. The RuntimeClass admission controller will reject Pod create requests which have the overhead already set. If RuntimeClass is configured and selected in the PodSpec, Overhead will be set to the value defined in the corresponding RuntimeClass, otherwise it will remain unset and treated as zero. More info: https://git.k8s.io/enhancements/keps/sig-node/688-pod-overhead/README.md
     */
    overhead?: {
        [name: string]: Quantity;
    };
    /**
     * PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset.
     */
    preemptionPolicy?: string;
    /**
     * The priority value. Various system components use this field to find the priority of the pod. When Priority Admission Controller is enabled, it prevents users from setting this field. The admission controller populates this field from PriorityClassName. The higher the value, the higher the priority.
     */
    priority?: number;
    /**
     * If specified, indicates the pod's priority. "system-node-critical" and "system-cluster-critical" are two special keywords which indicate the highest priorities with the former being the highest priority. Any other name must be defined by creating a PriorityClass object with that name. If not specified, the pod priority will be default or zero if there is no default.
     */
    priorityClassName?: string;
    /**
     * If specified, all readiness gates will be evaluated for pod readiness. A pod is ready when all its containers are ready AND all conditions specified in the readiness gates have status equal to "True" More info: https://git.k8s.io/enhancements/keps/sig-network/580-pod-readiness-gates
     */
    readinessGates?: Array<PodReadinessGate>;
    /**
     * Restart policy for all containers within the pod. One of Always, OnFailure, Never. Default to Always. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy
     *
     *
     */
    restartPolicy?: string;
    /**
     * RuntimeClassName refers to a RuntimeClass object in the node.k8s.io group, which should be used to run this pod.  If no RuntimeClass resource matches the named class, the pod will not be run. If unset or empty, the "legacy" RuntimeClass will be used, which is an implicit class with an empty definition that uses the default runtime handler. More info: https://git.k8s.io/enhancements/keps/sig-node/585-runtime-class
     */
    runtimeClassName?: string;
    /**
     * If specified, the pod will be dispatched by specified scheduler. If not specified, the pod will be dispatched by default scheduler.
     */
    schedulerName?: string;
    /**
     * SecurityContext holds pod-level security attributes and common container settings. Optional: Defaults to empty.  See type description for default values of each field.
     */
    securityContext?: PodSecurityContext;
    /**
     * DeprecatedServiceAccount is a depreciated alias for ServiceAccountName. Deprecated: Use serviceAccountName instead.
     */
    serviceAccount?: string;
    /**
     * ServiceAccountName is the name of the ServiceAccount to use to run this pod. More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
     */
    serviceAccountName?: string;
    /**
     * If true the pod's hostname will be configured as the pod's FQDN, rather than the leaf name (the default). In Linux containers, this means setting the FQDN in the hostname field of the kernel (the nodename field of struct utsname). In Windows containers, this means setting the registry value of hostname for the registry key HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Services\Tcpip\Parameters to FQDN. If a pod does not have FQDN, this has no effect. Default to false.
     */
    setHostnameAsFQDN?: boolean;
    /**
     * Share a single process namespace between all of the containers in a pod. When this is set containers will be able to view and signal processes from other containers in the same pod, and the first process in each container will not be assigned PID 1. HostPID and ShareProcessNamespace cannot both be set. Optional: Default to false.
     */
    shareProcessNamespace?: boolean;
    /**
     * If specified, the fully qualified Pod hostname will be "<hostname>.<subdomain>.<pod namespace>.svc.<cluster domain>". If not specified, the pod will not have a domainname at all.
     */
    subdomain?: string;
    /**
     * Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete request. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). If this value is nil, the default grace period will be used instead. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. Defaults to 30 seconds.
     */
    terminationGracePeriodSeconds?: number;
    /**
     * If specified, the pod's tolerations.
     */
    tolerations?: Array<Toleration>;
    /**
     * TopologySpreadConstraints describes how a group of pods ought to spread across topology domains. Scheduler will schedule pods in a way which abides by the constraints. All topologySpreadConstraints are ANDed.
     */
    topologySpreadConstraints?: Array<TopologySpreadConstraint>;
    /**
     * List of volumes that can be mounted by containers belonging to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes
     */
    volumes?: Array<Volume>;
}
/**
 * PodStatus represents information about the status of a pod. Status may trail the actual state of a system, especially if the node that hosts the pod cannot contact the control plane.
 */
export type PodStatus = {
    /**
     * Current service state of pod. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions
     */
    conditions?: Array<PodCondition>;
    /**
     * The list has one entry per container in the manifest. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-and-container-status
     */
    containerStatuses?: Array<ContainerStatus>;
    /**
     * Status for any ephemeral containers that have run in this pod. This field is beta-level and available on clusters that haven't disabled the EphemeralContainers feature gate.
     */
    ephemeralContainerStatuses?: Array<ContainerStatus>;
    /**
     * IP address of the host to which the pod is assigned. Empty if not yet scheduled.
     */
    hostIP?: string;
    /**
     * The list has one entry per init container in the manifest. The most recent successful init container will have ready = true, the most recently started container will have startTime set. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-and-container-status
     */
    initContainerStatuses?: Array<ContainerStatus>;
    /**
     * A human readable message indicating details about why the pod is in this condition.
     */
    message?: string;
    /**
     * nominatedNodeName is set only when this pod preempts other pods on the node, but it cannot be scheduled right away as preemption victims receive their graceful termination periods. This field does not guarantee that the pod will be scheduled on this node. Scheduler may decide to place the pod elsewhere if other nodes become available sooner. Scheduler may also decide to give the resources on this node to a higher priority pod that is created after preemption. As a result, this field may be different than PodSpec.nodeName when the pod is scheduled.
     */
    nominatedNodeName?: string;
    /**
     * The phase of a Pod is a simple, high-level summary of where the Pod is in its lifecycle. The conditions array, the reason and message fields, and the individual container status arrays contain more detail about the pod's status. There are five possible phase values:
     *
     * Pending: The pod has been accepted by the Kubernetes system, but one or more of the container images has not been created. This includes time before being scheduled as well as time spent downloading images over the network, which could take a while. Running: The pod has been bound to a node, and all of the containers have been created. At least one container is still running, or is in the process of starting or restarting. Succeeded: All containers in the pod have terminated in success, and will not be restarted. Failed: All containers in the pod have terminated, and at least one container has terminated in failure. The container either exited with non-zero status or was terminated by the system. Unknown: For some reason the state of the pod could not be obtained, typically due to an error in communicating with the host of the pod.
     *
     * More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-phase
     *
     *
     */
    phase?: string;
    /**
     * IP address allocated to the pod. Routable at least within the cluster. Empty if not yet allocated.
     */
    podIP?: string;
    /**
     * podIPs holds the IP addresses allocated to the pod. If this field is specified, the 0th entry must match the podIP field. Pods may be allocated at most 1 value for each of IPv4 and IPv6. This list is empty if no IPs have been allocated yet.
     */
    podIPs?: Array<PodIP>;
    /**
     * The Quality of Service (QOS) classification assigned to the pod based on resource requirements See PodQOSClass type for available QOS classes More info: https://git.k8s.io/community/contributors/design-proposals/node/resource-qos.md
     *
     *
     */
    qosClass?: string;
    /**
     * A brief CamelCase message indicating details about why the pod is in this state. e.g. 'Evicted'
     */
    reason?: string;
    /**
     * RFC 3339 date and time at which the object was acknowledged by the Kubelet. This is before the Kubelet pulled the container image(s) for the pod.
     */
    startTime?: Time;
}
/**
 * PodTemplate describes a template for creating copies of a predefined pod.
 */
export type PodTemplate = {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "v1";
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "PodTemplate";
    /**
     * Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     */
    metadata?: ObjectMeta;
    /**
     * Template defines the pods that will be created from this pod template. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     */
    template?: PodTemplateSpec;
}
/**
 * PodTemplateList is a list of PodTemplates.
 */
export type PodTemplateList = {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "v1";
    /**
     * List of pod templates
     */
    items: Array<PodTemplate>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "PodTemplateList";
    /**
     * Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    metadata?: ListMeta;
}
/**
 * PodTemplateSpec describes the data a pod should have when created from a template
 */
export type PodTemplateSpec = {
    /**
     * Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     */
    metadata?: ObjectMeta;
    /**
     * Specification of the desired behavior of the pod. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     */
    spec?: PodSpec;
}
export type PortStatus = {
    /**
     * Error is to record the problem with the service port The format of the error shall comply with the following rules: - built-in error values shall be specified in this file and those shall use
     *   CamelCase names
     * - cloud provider specific error values must have names that comply with the
     *   format foo.example.com/CamelCase.
     */
    error?: string;
    /**
     * Port is the port number of the service port of which status is recorded here
     */
    port: number;
    /**
     * Protocol is the protocol of the service port of which status is recorded here The supported values are: "TCP", "UDP", "SCTP"
     *
     *
     */
    protocol: string;
}
/**
 * PortworxVolumeSource represents a Portworx volume resource.
 */
export type PortworxVolumeSource = {
    /**
     * fSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs". Implicitly inferred to be "ext4" if unspecified.
     */
    fsType?: string;
    /**
     * readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
     */
    readOnly?: boolean;
    /**
     * volumeID uniquely identifies a Portworx volume
     */
    volumeID: string;
}
/**
 * An empty preferred scheduling term matches all objects with implicit weight 0 (i.e. it's a no-op). A null preferred scheduling term matches no objects (i.e. is also a no-op).
 */
export type PreferredSchedulingTerm = {
    /**
     * A node selector term, associated with the corresponding weight.
     */
    preference: NodeSelectorTerm;
    /**
     * Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.
     */
    weight: number;
}
/**
 * Probe describes a health check to be performed against a container to determine whether it is alive or ready to receive traffic.
 */
export type Probe = {
    /**
     * Exec specifies the action to take.
     */
    exec?: ExecAction;
    /**
     * Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.
     */
    failureThreshold?: number;
    /**
     * GRPC specifies an action involving a GRPC port. This is a beta field and requires enabling GRPCContainerProbe feature gate.
     */
    grpc?: GRPCAction;
    /**
     * HTTPGet specifies the http request to perform.
     */
    httpGet?: HTTPGetAction;
    /**
     * Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
     */
    initialDelaySeconds?: number;
    /**
     * How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.
     */
    periodSeconds?: number;
    /**
     * Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.
     */
    successThreshold?: number;
    /**
     * TCPSocket specifies an action involving a TCP port.
     */
    tcpSocket?: TCPSocketAction;
    /**
     * Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.
     */
    terminationGracePeriodSeconds?: number;
    /**
     * Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
     */
    timeoutSeconds?: number;
}
/**
 * Represents a projected volume source
 */
export type ProjectedVolumeSource = {
    /**
     * defaultMode are the mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
     */
    defaultMode?: number;
    /**
     * sources is the list of volume projections
     */
    sources?: Array<VolumeProjection>;
}
/**
 * Represents a Quobyte mount that lasts the lifetime of a pod. Quobyte volumes do not support ownership management or SELinux relabeling.
 */
export type QuobyteVolumeSource = {
    /**
     * group to map volume access to Default is no group
     */
    group?: string;
    /**
     * readOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.
     */
    readOnly?: boolean;
    /**
     * registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes
     */
    registry: string;
    /**
     * tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin
     */
    tenant?: string;
    /**
     * user to map volume access to Defaults to serivceaccount user
     */
    user?: string;
    /**
     * volume is a string that references an already created Quobyte volume by name.
     */
    volume: string;
}
/**
 * Represents a Rados Block Device mount that lasts the lifetime of a pod. RBD volumes support ownership management and SELinux relabeling.
 */
export type RBDPersistentVolumeSource = {
    /**
     * fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd
     */
    fsType?: string;
    /**
     * image is the rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
     */
    image: string;
    /**
     * keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
     */
    keyring?: string;
    /**
     * monitors is a collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
     */
    monitors: Array<string>;
    /**
     * pool is the rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
     */
    pool?: string;
    /**
     * readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
     */
    readOnly?: boolean;
    /**
     * secretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
     */
    secretRef?: SecretReference;
    /**
     * user is the rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
     */
    user?: string;
}
/**
 * Represents a Rados Block Device mount that lasts the lifetime of a pod. RBD volumes support ownership management and SELinux relabeling.
 */
export type RBDVolumeSource = {
    /**
     * fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd
     */
    fsType?: string;
    /**
     * image is the rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
     */
    image: string;
    /**
     * keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
     */
    keyring?: string;
    /**
     * monitors is a collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
     */
    monitors: Array<string>;
    /**
     * pool is the rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
     */
    pool?: string;
    /**
     * readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
     */
    readOnly?: boolean;
    /**
     * secretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
     */
    secretRef?: LocalObjectReference;
    /**
     * user is the rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
     */
    user?: string;
}
/**
 * ReplicationController represents the configuration of a replication controller.
 */
export type ReplicationController = {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "v1";
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "ReplicationController";
    /**
     * If the Labels of a ReplicationController are empty, they are defaulted to be the same as the Pod(s) that the replication controller manages. Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     */
    metadata?: ObjectMeta;
    /**
     * Spec defines the specification of the desired behavior of the replication controller. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     */
    spec?: ReplicationControllerSpec;
    /**
     * Status is the most recently observed status of the replication controller. This data may be out of date by some window of time. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     */
    readonly status?: ReplicationControllerStatus;
}
/**
 * ReplicationControllerCondition describes the state of a replication controller at a certain point.
 */
export type ReplicationControllerCondition = {
    /**
     * The last time the condition transitioned from one status to another.
     */
    lastTransitionTime?: Time;
    /**
     * A human readable message indicating details about the transition.
     */
    message?: string;
    /**
     * The reason for the condition's last transition.
     */
    reason?: string;
    /**
     * Status of the condition, one of True, False, Unknown.
     */
    status: string;
    /**
     * Type of replication controller condition.
     */
    type: string;
}
/**
 * ReplicationControllerList is a collection of replication controllers.
 */
export type ReplicationControllerList = {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "v1";
    /**
     * List of replication controllers. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller
     */
    items: Array<ReplicationController>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "ReplicationControllerList";
    /**
     * Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    metadata?: ListMeta;
}
/**
 * ReplicationControllerSpec is the specification of a replication controller.
 */
export type ReplicationControllerSpec = {
    /**
     * Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)
     */
    minReadySeconds?: number;
    /**
     * Replicas is the number of desired replicas. This is a pointer to distinguish between explicit zero and unspecified. Defaults to 1. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#what-is-a-replicationcontroller
     */
    replicas?: number;
    /**
     * Selector is a label query over pods that should match the Replicas count. If Selector is empty, it is defaulted to the labels present on the Pod template. Label keys and values that must match in order to be controlled by this replication controller, if empty defaulted to labels on Pod template. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
     */
    selector?: {
        [name: string]: string;
    };
    /**
     * Template is the object that describes the pod that will be created if insufficient replicas are detected. This takes precedence over a TemplateRef. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#pod-template
     */
    template?: PodTemplateSpec;
}
/**
 * ReplicationControllerStatus represents the current status of a replication controller.
 */
export type ReplicationControllerStatus = {
    /**
     * The number of available replicas (ready for at least minReadySeconds) for this replication controller.
     */
    availableReplicas?: number;
    /**
     * Represents the latest available observations of a replication controller's current state.
     */
    conditions?: Array<ReplicationControllerCondition>;
    /**
     * The number of pods that have labels matching the labels of the pod template of the replication controller.
     */
    fullyLabeledReplicas?: number;
    /**
     * ObservedGeneration reflects the generation of the most recently observed replication controller.
     */
    observedGeneration?: number;
    /**
     * The number of ready replicas for this replication controller.
     */
    readyReplicas?: number;
    /**
     * Replicas is the most recently oberved number of replicas. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#what-is-a-replicationcontroller
     */
    replicas: number;
}
/**
 * ResourceFieldSelector represents container resources (cpu, memory) and their output format
 */
export type ResourceFieldSelector = {
    /**
     * Container name: required for volumes, optional for env vars
     */
    containerName?: string;
    /**
     * Specifies the output format of the exposed resources, defaults to "1"
     */
    divisor?: Quantity;
    /**
     * Required: resource to select
     */
    resource: string;
}
/**
 * ResourceQuota sets aggregate quota restrictions enforced per namespace
 */
export type ResourceQuota = {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "v1";
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "ResourceQuota";
    /**
     * Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     */
    metadata?: ObjectMeta;
    /**
     * Spec defines the desired quota. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     */
    spec?: ResourceQuotaSpec;
    /**
     * Status defines the actual enforced quota and its current usage. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     */
    status?: ResourceQuotaStatus;
}
/**
 * ResourceQuotaList is a list of ResourceQuota items.
 */
export type ResourceQuotaList = {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "v1";
    /**
     * Items is a list of ResourceQuota objects. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/
     */
    items: Array<ResourceQuota>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "ResourceQuotaList";
    /**
     * Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    metadata?: ListMeta;
}
/**
 * ResourceQuotaSpec defines the desired hard limits to enforce for Quota.
 */
export type ResourceQuotaSpec = {
    /**
     * hard is the set of desired hard limits for each named resource. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/
     */
    hard?: {
        [name: string]: Quantity;
    };
    /**
     * scopeSelector is also a collection of filters like scopes that must match each object tracked by a quota but expressed using ScopeSelectorOperator in combination with possible values. For a resource to match, both scopes AND scopeSelector (if specified in spec), must be matched.
     */
    scopeSelector?: ScopeSelector;
    /**
     * A collection of filters that must match each object tracked by a quota. If not specified, the quota matches all objects.
     */
    scopes?: Array<string>;
}
/**
 * ResourceQuotaStatus defines the enforced hard limits and observed use.
 */
export type ResourceQuotaStatus = {
    /**
     * Hard is the set of enforced hard limits for each named resource. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/
     */
    hard?: {
        [name: string]: Quantity;
    };
    /**
     * Used is the current observed total usage of the resource in the namespace.
     */
    used?: {
        [name: string]: Quantity;
    };
}
/**
 * ResourceRequirements describes the compute resource requirements.
 */
export type ResourceRequirements = {
    /**
     * Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
     */
    limits?: {
        [name: string]: Quantity;
    };
    /**
     * Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
     */
    requests?: {
        [name: string]: Quantity;
    };
}
/**
 * SELinuxOptions are the labels to be applied to the container
 */
export type SELinuxOptions = {
    /**
     * Level is SELinux level label that applies to the container.
     */
    level?: string;
    /**
     * Role is a SELinux role label that applies to the container.
     */
    role?: string;
    /**
     * Type is a SELinux type label that applies to the container.
     */
    type?: string;
    /**
     * User is a SELinux user label that applies to the container.
     */
    user?: string;
}
/**
 * ScaleIOPersistentVolumeSource represents a persistent ScaleIO volume
 */
export type ScaleIOPersistentVolumeSource = {
    /**
     * fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Default is "xfs"
     */
    fsType?: string;
    /**
     * gateway is the host address of the ScaleIO API Gateway.
     */
    gateway: string;
    /**
     * protectionDomain is the name of the ScaleIO Protection Domain for the configured storage.
     */
    protectionDomain?: string;
    /**
     * readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
     */
    readOnly?: boolean;
    /**
     * secretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.
     */
    secretRef: SecretReference;
    /**
     * sslEnabled is the flag to enable/disable SSL communication with Gateway, default false
     */
    sslEnabled?: boolean;
    /**
     * storageMode indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.
     */
    storageMode?: string;
    /**
     * storagePool is the ScaleIO Storage Pool associated with the protection domain.
     */
    storagePool?: string;
    /**
     * system is the name of the storage system as configured in ScaleIO.
     */
    system: string;
    /**
     * volumeName is the name of a volume already created in the ScaleIO system that is associated with this volume source.
     */
    volumeName?: string;
}
/**
 * ScaleIOVolumeSource represents a persistent ScaleIO volume
 */
export type ScaleIOVolumeSource = {
    /**
     * fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Default is "xfs".
     */
    fsType?: string;
    /**
     * gateway is the host address of the ScaleIO API Gateway.
     */
    gateway: string;
    /**
     * protectionDomain is the name of the ScaleIO Protection Domain for the configured storage.
     */
    protectionDomain?: string;
    /**
     * readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
     */
    readOnly?: boolean;
    /**
     * secretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.
     */
    secretRef: LocalObjectReference;
    /**
     * sslEnabled Flag enable/disable SSL communication with Gateway, default false
     */
    sslEnabled?: boolean;
    /**
     * storageMode indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.
     */
    storageMode?: string;
    /**
     * storagePool is the ScaleIO Storage Pool associated with the protection domain.
     */
    storagePool?: string;
    /**
     * system is the name of the storage system as configured in ScaleIO.
     */
    system: string;
    /**
     * volumeName is the name of a volume already created in the ScaleIO system that is associated with this volume source.
     */
    volumeName?: string;
}
/**
 * A scope selector represents the AND of the selectors represented by the scoped-resource selector requirements.
 */
export type ScopeSelector = {
    /**
     * A list of scope selector requirements by scope of the resources.
     */
    matchExpressions?: Array<ScopedResourceSelectorRequirement>;
}
/**
 * A scoped-resource selector requirement is a selector that contains values, a scope name, and an operator that relates the scope name and values.
 */
export type ScopedResourceSelectorRequirement = {
    /**
     * Represents a scope's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist.
     *
     *
     */
    operator: string;
    /**
     * The name of the scope that the selector applies to.
     *
     *
     */
    scopeName: string;
    /**
     * An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
     */
    values?: Array<string>;
}
/**
 * SeccompProfile defines a pod/container's seccomp profile settings. Only one profile source may be set.
 */
export type SeccompProfile = {
    /**
     * localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must only be set if type is "Localhost".
     */
    localhostProfile?: string;
    /**
     * type indicates which kind of seccomp profile will be applied. Valid options are:
     *
     * Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied.
     *
     *
     */
    type: string;
}
/**
 * Secret holds secret data of a certain type. The total bytes of the values in the Data field must be less than MaxSecretSize bytes.
 */
export type Secret = {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "v1";
    /**
     * Data contains the secret data. Each key must consist of alphanumeric characters, '-', '_' or '.'. The serialized form of the secret data is a base64 encoded string, representing the arbitrary (possibly non-string) data value here. Described in https://tools.ietf.org/html/rfc4648#section-4
     */
    data?: {
        [name: string]: string;
    };
    /**
     * Immutable, if set to true, ensures that data stored in the Secret cannot be updated (only object metadata can be modified). If not set to true, the field can be modified at any time. Defaulted to nil.
     */
    immutable?: boolean;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "Secret";
    /**
     * Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     */
    metadata?: ObjectMeta;
    /**
     * stringData allows specifying non-binary secret data in string form. It is provided as a write-only input field for convenience. All keys and values are merged into the data field on write, overwriting any existing values. The stringData field is never output when reading from the API.
     */
    stringData?: {
        [name: string]: string;
    };
    /**
     * Used to facilitate programmatic handling of secret data. More info: https://kubernetes.io/docs/concepts/configuration/secret/#secret-types
     */
    type?: string;
}
/**
 * SecretEnvSource selects a Secret to populate the environment variables with.
 *
 * The contents of the target Secret's Data field will represent the key-value pairs as environment variables.
 */
export type SecretEnvSource = {
    /**
     * Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
     */
    name?: string;
    /**
     * Specify whether the Secret must be defined
     */
    optional?: boolean;
}
/**
 * SecretKeySelector selects a key of a Secret.
 */
export type SecretKeySelector = {
    /**
     * The key of the secret to select from.  Must be a valid secret key.
     */
    key: string;
    /**
     * Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
     */
    name?: string;
    /**
     * Specify whether the Secret or its key must be defined
     */
    optional?: boolean;
}
/**
 * SecretList is a list of Secret.
 */
export type SecretList = {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "v1";
    /**
     * Items is a list of secret objects. More info: https://kubernetes.io/docs/concepts/configuration/secret
     */
    items: Array<Secret>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "SecretList";
    /**
     * Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    metadata?: ListMeta;
}
/**
 * Adapts a secret into a projected volume.
 *
 * The contents of the target Secret's Data field will be presented in a projected volume as files using the keys in the Data field as the file names. Note that this is identical to a secret volume source without the default mode.
 */
export type SecretProjection = {
    /**
     * items if unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.
     */
    items?: Array<KeyToPath>;
    /**
     * Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
     */
    name?: string;
    /**
     * optional field specify whether the Secret or its key must be defined
     */
    optional?: boolean;
}
/**
 * SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace
 */
export type SecretReference = {
    /**
     * name is unique within a namespace to reference a secret resource.
     */
    name?: string;
    /**
     * namespace defines the space within which the secret name must be unique.
     */
    namespace?: string;
}
/**
 * Adapts a Secret into a volume.
 *
 * The contents of the target Secret's Data field will be presented in a volume as files using the keys in the Data field as the file names. Secret volumes support ownership management and SELinux relabeling.
 */
export type SecretVolumeSource = {
    /**
     * defaultMode is Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
     */
    defaultMode?: number;
    /**
     * items If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.
     */
    items?: Array<KeyToPath>;
    /**
     * optional field specify whether the Secret or its keys must be defined
     */
    optional?: boolean;
    /**
     * secretName is the name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret
     */
    secretName?: string;
}
/**
 * SecurityContext holds security configuration that will be applied to a container. Some fields are present in both SecurityContext and PodSecurityContext.  When both are set, the values in SecurityContext take precedence.
 */
export type SecurityContext = {
    /**
     * AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN Note that this field cannot be set when spec.os.name is windows.
     */
    allowPrivilegeEscalation?: boolean;
    /**
     * The capabilities to add/drop when running containers. Defaults to the default set of capabilities granted by the container runtime. Note that this field cannot be set when spec.os.name is windows.
     */
    capabilities?: Capabilities;
    /**
     * Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false. Note that this field cannot be set when spec.os.name is windows.
     */
    privileged?: boolean;
    /**
     * procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths. This requires the ProcMountType feature flag to be enabled. Note that this field cannot be set when spec.os.name is windows.
     */
    procMount?: string;
    /**
     * Whether this container has a read-only root filesystem. Default is false. Note that this field cannot be set when spec.os.name is windows.
     */
    readOnlyRootFilesystem?: boolean;
    /**
     * The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.
     */
    runAsGroup?: number;
    /**
     * Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
     */
    runAsNonRoot?: boolean;
    /**
     * The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.
     */
    runAsUser?: number;
    /**
     * The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.
     */
    seLinuxOptions?: SELinuxOptions;
    /**
     * The seccomp options to use by this container. If seccomp options are provided at both the pod & container level, the container options override the pod options. Note that this field cannot be set when spec.os.name is windows.
     */
    seccompProfile?: SeccompProfile;
    /**
     * The Windows specific settings applied to all containers. If unspecified, the options from the PodSecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.
     */
    windowsOptions?: WindowsSecurityContextOptions;
}
/**
 * Service is a named abstraction of software service (for example, mysql) consisting of local port (for example 3306) that the proxy listens on, and the selector that determines which pods will answer requests sent through the proxy.
 */
export type Service = {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "v1";
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "Service";
    /**
     * Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     */
    metadata?: ObjectMeta;
    /**
     * Spec defines the behavior of a service. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     */
    spec?: ServiceSpec;
    /**
     * Most recently observed status of the service. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     */
    readonly status?: ServiceStatus;
}
/**
 * ServiceAccount binds together: * a name, understood by users, and perhaps by peripheral systems, for an identity * a principal that can be authenticated and authorized * a set of secrets
 */
export type ServiceAccount = {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "v1";
    /**
     * AutomountServiceAccountToken indicates whether pods running as this service account should have an API token automatically mounted. Can be overridden at the pod level.
     */
    automountServiceAccountToken?: boolean;
    /**
     * ImagePullSecrets is a list of references to secrets in the same namespace to use for pulling any images in pods that reference this ServiceAccount. ImagePullSecrets are distinct from Secrets because Secrets can be mounted in the pod, but ImagePullSecrets are only accessed by the kubelet. More info: https://kubernetes.io/docs/concepts/containers/images/#specifying-imagepullsecrets-on-a-pod
     */
    imagePullSecrets?: Array<LocalObjectReference>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "ServiceAccount";
    /**
     * Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     */
    metadata?: ObjectMeta;
    /**
     * Secrets is a list of the secrets in the same namespace that pods running using this ServiceAccount are allowed to use. Pods are only limited to this list if this service account has a "kubernetes.io/enforce-mountable-secrets" annotation set to "true". This field should not be used to find auto-generated service account token secrets for use outside of pods. Instead, tokens can be requested directly using the TokenRequest API, or service account token secrets can be manually created. More info: https://kubernetes.io/docs/concepts/configuration/secret
     */
    secrets?: Array<ObjectReference>;
}
/**
 * ServiceAccountList is a list of ServiceAccount objects
 */
export type ServiceAccountList = {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "v1";
    /**
     * List of ServiceAccounts. More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
     */
    items: Array<ServiceAccount>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "ServiceAccountList";
    /**
     * Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    metadata?: ListMeta;
}
/**
 * ServiceAccountTokenProjection represents a projected service account token volume. This projection can be used to insert a service account token into the pods runtime filesystem for use against APIs (Kubernetes API Server or otherwise).
 */
export type ServiceAccountTokenProjection = {
    /**
     * audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token. The audience defaults to the identifier of the apiserver.
     */
    audience?: string;
    /**
     * expirationSeconds is the requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token. The kubelet will start trying to rotate the token if the token is older than 80 percent of its time to live or if the token is older than 24 hours.Defaults to 1 hour and must be at least 10 minutes.
     */
    expirationSeconds?: number;
    /**
     * path is the path relative to the mount point of the file to project the token into.
     */
    path: string;
}
/**
 * ServiceList holds a list of services.
 */
export type ServiceList = {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "v1";
    /**
     * List of services
     */
    items: Array<Service>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "ServiceList";
    /**
     * Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    metadata?: ListMeta;
}
/**
 * ServicePort contains information on service's port.
 */
export type ServicePort = {
    /**
     * The application protocol for this port. This field follows standard Kubernetes label syntax. Un-prefixed names are reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names). Non-standard protocols should use prefixed names such as mycompany.com/my-custom-protocol.
     */
    appProtocol?: string;
    /**
     * The name of this port within the service. This must be a DNS_LABEL. All ports within a ServiceSpec must have unique names. When considering the endpoints for a Service, this must match the 'name' field in the EndpointPort. Optional if only one ServicePort is defined on this service.
     */
    name?: string;
    /**
     * The port on each node on which this service is exposed when type is NodePort or LoadBalancer.  Usually assigned by the system. If a value is specified, in-range, and not in use it will be used, otherwise the operation will fail.  If not specified, a port will be allocated if this Service requires one.  If this field is specified when creating a Service which does not need it, creation will fail. This field will be wiped when updating a Service to no longer need it (e.g. changing type from NodePort to ClusterIP). More info: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport
     */
    nodePort?: number;
    /**
     * The port that will be exposed by this service.
     */
    port: number;
    /**
     * The IP protocol for this port. Supports "TCP", "UDP", and "SCTP". Default is TCP.
     *
     *
     */
    protocol?: string;
    /**
     * Number or name of the port to access on the pods targeted by the service. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME. If this is a string, it will be looked up as a named port in the target Pod's container ports. If this is not specified, the value of the 'port' field is used (an identity map). This field is ignored for services with clusterIP=None, and should be omitted or set equal to the 'port' field. More info: https://kubernetes.io/docs/concepts/services-networking/service/#defining-a-service
     */
    targetPort?: number | string;
}
/**
 * ServiceSpec describes the attributes that a user creates on a service.
 */
export type ServiceSpec = {
    /**
     * allocateLoadBalancerNodePorts defines if NodePorts will be automatically allocated for services with type LoadBalancer.  Default is "true". It may be set to "false" if the cluster load-balancer does not rely on NodePorts.  If the caller requests specific NodePorts (by specifying a value), those requests will be respected, regardless of this field. This field may only be set for services with type LoadBalancer and will be cleared if the type is changed to any other type.
     */
    allocateLoadBalancerNodePorts?: boolean;
    /**
     * clusterIP is the IP address of the service and is usually assigned randomly. If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be blank) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above).  Valid values are "None", empty string (""), or a valid IP address. Setting this to "None" makes a "headless service" (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required.  Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a Service of type ExternalName, creation will fail. This field will be wiped when updating a Service to type ExternalName. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
     */
    clusterIP?: string;
    /**
     * ClusterIPs is a list of IP addresses assigned to this service, and are usually assigned randomly.  If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be empty) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above).  Valid values are "None", empty string (""), or a valid IP address.  Setting this to "None" makes a "headless service" (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required.  Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a Service of type ExternalName, creation will fail. This field will be wiped when updating a Service to type ExternalName.  If this field is not specified, it will be initialized from the clusterIP field.  If this field is specified, clients must ensure that clusterIPs[0] and clusterIP have the same value.
     *
     * This field may hold a maximum of two entries (dual-stack IPs, in either order). These IPs must correspond to the values of the ipFamilies field. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
     */
    clusterIPs?: Array<string>;
    /**
     * externalIPs is a list of IP addresses for which nodes in the cluster will also accept traffic for this service.  These IPs are not managed by Kubernetes.  The user is responsible for ensuring that traffic arrives at a node with this IP.  A common example is external load-balancers that are not part of the Kubernetes system.
     */
    externalIPs?: Array<string>;
    /**
     * externalName is the external reference that discovery mechanisms will return as an alias for this service (e.g. a DNS CNAME record). No proxying will be involved.  Must be a lowercase RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) and requires `type` to be "ExternalName".
     */
    externalName?: string;
    /**
     * externalTrafficPolicy describes how nodes distribute service traffic they receive on one of the Service's "externally-facing" addresses (NodePorts, ExternalIPs, and LoadBalancer IPs). If set to "Local", the proxy will configure the service in a way that assumes that external load balancers will take care of balancing the service traffic between nodes, and so each node will deliver traffic only to the node-local endpoints of the service, without masquerading the client source IP. (Traffic mistakenly sent to a node with no endpoints will be dropped.) The default value, "Cluster", uses the standard behavior of routing to all endpoints evenly (possibly modified by topology and other features). Note that traffic sent to an External IP or LoadBalancer IP from within the cluster will always get "Cluster" semantics, but clients sending to a NodePort from within the cluster may need to take traffic policy into account when picking a node.
     *
     *
     */
    externalTrafficPolicy?: string;
    /**
     * healthCheckNodePort specifies the healthcheck nodePort for the service. This only applies when type is set to LoadBalancer and externalTrafficPolicy is set to Local. If a value is specified, is in-range, and is not in use, it will be used.  If not specified, a value will be automatically allocated.  External systems (e.g. load-balancers) can use this port to determine if a given node holds endpoints for this service or not.  If this field is specified when creating a Service which does not need it, creation will fail. This field will be wiped when updating a Service to no longer need it (e.g. changing type).
     */
    healthCheckNodePort?: number;
    /**
     * InternalTrafficPolicy describes how nodes distribute service traffic they receive on the ClusterIP. If set to "Local", the proxy will assume that pods only want to talk to endpoints of the service on the same node as the pod, dropping the traffic if there are no local endpoints. The default value, "Cluster", uses the standard behavior of routing to all endpoints evenly (possibly modified by topology and other features).
     */
    internalTrafficPolicy?: string;
    /**
     * IPFamilies is a list of IP families (e.g. IPv4, IPv6) assigned to this service. This field is usually assigned automatically based on cluster configuration and the ipFamilyPolicy field. If this field is specified manually, the requested family is available in the cluster, and ipFamilyPolicy allows it, it will be used; otherwise creation of the service will fail. This field is conditionally mutable: it allows for adding or removing a secondary IP family, but it does not allow changing the primary IP family of the Service. Valid values are "IPv4" and "IPv6".  This field only applies to Services of types ClusterIP, NodePort, and LoadBalancer, and does apply to "headless" services. This field will be wiped when updating a Service to type ExternalName.
     *
     * This field may hold a maximum of two entries (dual-stack families, in either order).  These families must correspond to the values of the clusterIPs field, if specified. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field.
     */
    ipFamilies?: Array<string>;
    /**
     * IPFamilyPolicy represents the dual-stack-ness requested or required by this Service. If there is no value provided, then this field will be set to SingleStack. Services can be "SingleStack" (a single IP family), "PreferDualStack" (two IP families on dual-stack configured clusters or a single IP family on single-stack clusters), or "RequireDualStack" (two IP families on dual-stack configured clusters, otherwise fail). The ipFamilies and clusterIPs fields depend on the value of this field. This field will be wiped when updating a service to type ExternalName.
     */
    ipFamilyPolicy?: string;
    /**
     * loadBalancerClass is the class of the load balancer implementation this Service belongs to. If specified, the value of this field must be a label-style identifier, with an optional prefix, e.g. "internal-vip" or "example.com/internal-vip". Unprefixed names are reserved for end-users. This field can only be set when the Service type is 'LoadBalancer'. If not set, the default load balancer implementation is used, today this is typically done through the cloud provider integration, but should apply for any default implementation. If set, it is assumed that a load balancer implementation is watching for Services with a matching class. Any default load balancer implementation (e.g. cloud providers) should ignore Services that set this field. This field can only be set when creating or updating a Service to type 'LoadBalancer'. Once set, it can not be changed. This field will be wiped when a service is updated to a non 'LoadBalancer' type.
     */
    loadBalancerClass?: string;
    /**
     * Only applies to Service Type: LoadBalancer. This feature depends on whether the underlying cloud-provider supports specifying the loadBalancerIP when a load balancer is created. This field will be ignored if the cloud-provider does not support the feature. Deprecated: This field was under-specified and its meaning varies across implementations, and it cannot support dual-stack. As of Kubernetes v1.24, users are encouraged to use implementation-specific annotations when available. This field may be removed in a future API version.
     */
    loadBalancerIP?: string;
    /**
     * If specified and supported by the platform, this will restrict traffic through the cloud-provider load-balancer will be restricted to the specified client IPs. This field will be ignored if the cloud-provider does not support the feature." More info: https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/
     */
    loadBalancerSourceRanges?: Array<string>;
    /**
     * The list of ports that are exposed by this service. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
     */
    ports?: Array<ServicePort>;
    /**
     * publishNotReadyAddresses indicates that any agent which deals with endpoints for this Service should disregard any indications of ready/not-ready. The primary use case for setting this field is for a StatefulSet's Headless Service to propagate SRV DNS records for its Pods for the purpose of peer discovery. The Kubernetes controllers that generate Endpoints and EndpointSlice resources for Services interpret this to mean that all endpoints are considered "ready" even if the Pods themselves are not. Agents which consume only Kubernetes generated endpoints through the Endpoints or EndpointSlice resources can safely assume this behavior.
     */
    publishNotReadyAddresses?: boolean;
    /**
     * Route service traffic to pods with label keys and values matching this selector. If empty or not present, the service is assumed to have an external process managing its endpoints, which Kubernetes will not modify. Only applies to types ClusterIP, NodePort, and LoadBalancer. Ignored if type is ExternalName. More info: https://kubernetes.io/docs/concepts/services-networking/service/
     */
    selector?: {
        [name: string]: string;
    };
    /**
     * Supports "ClientIP" and "None". Used to maintain session affinity. Enable client IP based session affinity. Must be ClientIP or None. Defaults to None. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
     *
     *
     */
    sessionAffinity?: string;
    /**
     * sessionAffinityConfig contains the configurations of session affinity.
     */
    sessionAffinityConfig?: SessionAffinityConfig;
    /**
     * type determines how the Service is exposed. Defaults to ClusterIP. Valid options are ExternalName, ClusterIP, NodePort, and LoadBalancer. "ClusterIP" allocates a cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, by manual construction of an Endpoints object or EndpointSlice objects. If clusterIP is "None", no virtual IP is allocated and the endpoints are published as a set of endpoints rather than a virtual IP. "NodePort" builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the clusterIP. "LoadBalancer" builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the same endpoints as the clusterIP. "ExternalName" aliases this service to the specified externalName. Several other fields do not apply to ExternalName services. More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types
     *
     *
     */
    type?: string;
}
/**
 * ServiceStatus represents the current status of a service.
 */
export type ServiceStatus = {
    /**
     * Current service state
     */
    conditions?: Array<Condition>;
    /**
     * LoadBalancer contains the current status of the load-balancer, if one is present.
     */
    loadBalancer?: LoadBalancerStatus;
}
/**
 * SessionAffinityConfig represents the configurations of session affinity.
 */
export type SessionAffinityConfig = {
    /**
     * clientIP contains the configurations of Client IP based session affinity.
     */
    clientIP?: ClientIPConfig;
}
/**
 * Represents a StorageOS persistent volume resource.
 */
export type StorageOSPersistentVolumeSource = {
    /**
     * fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
     */
    fsType?: string;
    /**
     * readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
     */
    readOnly?: boolean;
    /**
     * secretRef specifies the secret to use for obtaining the StorageOS API credentials.  If not specified, default values will be attempted.
     */
    secretRef?: ObjectReference;
    /**
     * volumeName is the human-readable name of the StorageOS volume.  Volume names are only unique within a namespace.
     */
    volumeName?: string;
    /**
     * volumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod's namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to "default" if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.
     */
    volumeNamespace?: string;
}
/**
 * Represents a StorageOS persistent volume resource.
 */
export type StorageOSVolumeSource = {
    /**
     * fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
     */
    fsType?: string;
    /**
     * readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
     */
    readOnly?: boolean;
    /**
     * secretRef specifies the secret to use for obtaining the StorageOS API credentials.  If not specified, default values will be attempted.
     */
    secretRef?: LocalObjectReference;
    /**
     * volumeName is the human-readable name of the StorageOS volume.  Volume names are only unique within a namespace.
     */
    volumeName?: string;
    /**
     * volumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod's namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to "default" if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.
     */
    volumeNamespace?: string;
}
/**
 * Sysctl defines a kernel parameter to be set
 */
export type Sysctl = {
    /**
     * Name of a property to set
     */
    name: string;
    /**
     * Value of a property to set
     */
    value: string;
}
/**
 * TCPSocketAction describes an action based on opening a socket
 */
export type TCPSocketAction = {
    /**
     * Optional: Host name to connect to, defaults to the pod IP.
     */
    host?: string;
    /**
     * Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
     */
    port: number | string;
}
/**
 * The node this Taint is attached to has the "effect" on any pod that does not tolerate the Taint.
 */
export type Taint = {
    /**
     * Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.
     *
     *
     */
    effect: string;
    /**
     * Required. The taint key to be applied to a node.
     */
    key: string;
    /**
     * TimeAdded represents the time at which the taint was added. It is only written for NoExecute taints.
     */
    timeAdded?: Time;
    /**
     * The taint value corresponding to the taint key.
     */
    value?: string;
}
/**
 * The pod this Toleration is attached to tolerates any taint that matches the triple <key,value,effect> using the matching operator <operator>.
 */
export type Toleration = {
    /**
     * Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.
     *
     *
     */
    effect?: string;
    /**
     * Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.
     */
    key?: string;
    /**
     * Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.
     *
     *
     */
    operator?: string;
    /**
     * TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.
     */
    tolerationSeconds?: number;
    /**
     * Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.
     */
    value?: string;
}
/**
 * A topology selector requirement is a selector that matches given label. This is an alpha feature and may change in the future.
 */
export type TopologySelectorLabelRequirement = {
    /**
     * The label key that the selector applies to.
     */
    key: string;
    /**
     * An array of string values. One value must match the label to be selected. Each entry in Values is ORed.
     */
    values: Array<string>;
}
/**
 * A topology selector term represents the result of label queries. A null or empty topology selector term matches no objects. The requirements of them are ANDed. It provides a subset of functionality as NodeSelectorTerm. This is an alpha feature and may change in the future.
 */
export type TopologySelectorTerm = {
    /**
     * A list of topology selector requirements by labels.
     */
    matchLabelExpressions?: Array<TopologySelectorLabelRequirement>;
}
/**
 * TopologySpreadConstraint specifies how to spread matching pods among the given topology.
 */
export type TopologySpreadConstraint = {
    /**
     * LabelSelector is used to find matching pods. Pods that match this label selector are counted to determine the number of pods in their corresponding topology domain.
     */
    labelSelector?: LabelSelector;
    /**
     * MaxSkew describes the degree to which pods may be unevenly distributed. When `whenUnsatisfiable=DoNotSchedule`, it is the maximum permitted difference between the number of matching pods in the target topology and the global minimum. The global minimum is the minimum number of matching pods in an eligible domain or zero if the number of eligible domains is less than MinDomains. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 2/2/1: In this case, the global minimum is 1. | zone1 | zone2 | zone3 | |  P P  |  P P  |   P   | - if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 2/2/2; scheduling it onto zone1(zone2) would make the ActualSkew(3-1) on zone1(zone2) violate MaxSkew(1). - if MaxSkew is 2, incoming pod can be scheduled onto any zone. When `whenUnsatisfiable=ScheduleAnyway`, it is used to give higher precedence to topologies that satisfy it. It's a required field. Default value is 1 and 0 is not allowed.
     */
    maxSkew: number;
    /**
     * MinDomains indicates a minimum number of eligible domains. When the number of eligible domains with matching topology keys is less than minDomains, Pod Topology Spread treats "global minimum" as 0, and then the calculation of Skew is performed. And when the number of eligible domains with matching topology keys equals or greater than minDomains, this value has no effect on scheduling. As a result, when the number of eligible domains is less than minDomains, scheduler won't schedule more than maxSkew Pods to those domains. If value is nil, the constraint behaves as if MinDomains is equal to 1. Valid values are integers greater than 0. When value is not nil, WhenUnsatisfiable must be DoNotSchedule.
     *
     * For example, in a 3-zone cluster, MaxSkew is set to 2, MinDomains is set to 5 and pods with the same labelSelector spread as 2/2/2: | zone1 | zone2 | zone3 | |  P P  |  P P  |  P P  | The number of domains is less than 5(MinDomains), so "global minimum" is treated as 0. In this situation, new pod with the same labelSelector cannot be scheduled, because computed skew will be 3(3 - 0) if new Pod is scheduled to any of the three zones, it will violate MaxSkew.
     *
     * This is an alpha field and requires enabling MinDomainsInPodTopologySpread feature gate.
     */
    minDomains?: number;
    /**
     * NodeAffinityPolicy indicates how we will treat Pod's nodeAffinity/nodeSelector when calculating pod topology spread skew. Options are: - Honor: only nodes matching nodeAffinity/nodeSelector are included in the calculations. - Ignore: nodeAffinity/nodeSelector are ignored. All nodes are included in the calculations.
     *
     * If this value is nil, the behavior is equivalent to the Honor policy. This is a alpha-level feature enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.
     */
    nodeAffinityPolicy?: string;
    /**
     * NodeTaintsPolicy indicates how we will treat node taints when calculating pod topology spread skew. Options are: - Honor: nodes without taints, along with tainted nodes for which the incoming pod has a toleration, are included. - Ignore: node taints are ignored. All nodes are included.
     *
     * If this value is nil, the behavior is equivalent to the Ignore policy. This is a alpha-level feature enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.
     */
    nodeTaintsPolicy?: string;
    /**
     * TopologyKey is the key of node labels. Nodes that have a label with this key and identical values are considered to be in the same topology. We consider each <key, value> as a "bucket", and try to put balanced number of pods into each bucket. We define a domain as a particular instance of a topology. Also, we define an eligible domain as a domain whose nodes meet the requirements of nodeAffinityPolicy and nodeTaintsPolicy. e.g. If TopologyKey is "kubernetes.io/hostname", each Node is a domain of that topology. And, if TopologyKey is "topology.kubernetes.io/zone", each zone is a domain of that topology. It's a required field.
     */
    topologyKey: string;
    /**
     * WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfy the spread constraint. - DoNotSchedule (default) tells the scheduler not to schedule it. - ScheduleAnyway tells the scheduler to schedule the pod in any location,
     *   but giving higher precedence to topologies that would help reduce the
     *   skew.
     * A constraint is considered "Unsatisfiable" for an incoming pod if and only if every possible node assignment for that pod would violate "MaxSkew" on some topology. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 3/1/1: | zone1 | zone2 | zone3 | | P P P |   P   |   P   | If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduled to zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfies MaxSkew(1). In other words, the cluster can still be imbalanced, but scheduler won't make it *more* imbalanced. It's a required field.
     *
     *
     */
    whenUnsatisfiable: string;
}
/**
 * TypedLocalObjectReference contains enough information to let you locate the typed referenced object inside the same namespace.
 */
export type TypedLocalObjectReference = {
    /**
     * APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.
     */
    apiGroup?: string;
    /**
     * Kind is the type of resource being referenced
     */
    kind: string;
    /**
     * Name is the name of resource being referenced
     */
    name: string;
}
/**
 * Volume represents a named volume in a pod that may be accessed by any container in the pod.
 */
export type Volume = {
    /**
     * awsElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
     */
    awsElasticBlockStore?: AWSElasticBlockStoreVolumeSource;
    /**
     * azureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
     */
    azureDisk?: AzureDiskVolumeSource;
    /**
     * azureFile represents an Azure File Service mount on the host and bind mount to the pod.
     */
    azureFile?: AzureFileVolumeSource;
    /**
     * cephFS represents a Ceph FS mount on the host that shares a pod's lifetime
     */
    cephfs?: CephFSVolumeSource;
    /**
     * cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
     */
    cinder?: CinderVolumeSource;
    /**
     * configMap represents a configMap that should populate this volume
     */
    configMap?: ConfigMapVolumeSource;
    /**
     * csi (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).
     */
    csi?: CSIVolumeSource;
    /**
     * downwardAPI represents downward API about the pod that should populate this volume
     */
    downwardAPI?: DownwardAPIVolumeSource;
    /**
     * emptyDir represents a temporary directory that shares a pod's lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir
     */
    emptyDir?: EmptyDirVolumeSource;
    /**
     * ephemeral represents a volume that is handled by a cluster storage driver. The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed.
     *
     * Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity
     *    tracking are needed,
     * c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through
     *    a PersistentVolumeClaim (see EphemeralVolumeSource for more
     *    information on the connection between this volume type
     *    and PersistentVolumeClaim).
     *
     * Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod.
     *
     * Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information.
     *
     * A pod can use both types of ephemeral volumes and persistent volumes at the same time.
     */
    ephemeral?: EphemeralVolumeSource;
    /**
     * fc represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.
     */
    fc?: FCVolumeSource;
    /**
     * flexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.
     */
    flexVolume?: FlexVolumeSource;
    /**
     * flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running
     */
    flocker?: FlockerVolumeSource;
    /**
     * gcePersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
     */
    gcePersistentDisk?: GCEPersistentDiskVolumeSource;
    /**
     * gitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.
     */
    gitRepo?: GitRepoVolumeSource;
    /**
     * glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/glusterfs/README.md
     */
    glusterfs?: GlusterfsVolumeSource;
    /**
     * hostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
     */
    hostPath?: HostPathVolumeSource;
    /**
     * iscsi represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://examples.k8s.io/volumes/iscsi/README.md
     */
    iscsi?: ISCSIVolumeSource;
    /**
     * name of the volume. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
     */
    name: string;
    /**
     * nfs represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
     */
    nfs?: NFSVolumeSource;
    /**
     * persistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
     */
    persistentVolumeClaim?: PersistentVolumeClaimVolumeSource;
    /**
     * photonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine
     */
    photonPersistentDisk?: PhotonPersistentDiskVolumeSource;
    /**
     * portworxVolume represents a portworx volume attached and mounted on kubelets host machine
     */
    portworxVolume?: PortworxVolumeSource;
    /**
     * projected items for all in one resources secrets, configmaps, and downward API
     */
    projected?: ProjectedVolumeSource;
    /**
     * quobyte represents a Quobyte mount on the host that shares a pod's lifetime
     */
    quobyte?: QuobyteVolumeSource;
    /**
     * rbd represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md
     */
    rbd?: RBDVolumeSource;
    /**
     * scaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.
     */
    scaleIO?: ScaleIOVolumeSource;
    /**
     * secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret
     */
    secret?: SecretVolumeSource;
    /**
     * storageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.
     */
    storageos?: StorageOSVolumeSource;
    /**
     * vsphereVolume represents a vSphere volume attached and mounted on kubelets host machine
     */
    vsphereVolume?: VsphereVirtualDiskVolumeSource;
}
/**
 * volumeDevice describes a mapping of a raw block device within a container.
 */
export type VolumeDevice = {
    /**
     * devicePath is the path inside of the container that the device will be mapped to.
     */
    devicePath: string;
    /**
     * name must match the name of a persistentVolumeClaim in the pod
     */
    name: string;
}
/**
 * VolumeMount describes a mounting of a Volume within a container.
 */
export type VolumeMount = {
    /**
     * Path within the container at which the volume should be mounted.  Must not contain ':'.
     */
    mountPath: string;
    /**
     * mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.
     */
    mountPropagation?: string;
    /**
     * This must match the Name of a Volume.
     */
    name: string;
    /**
     * Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.
     */
    readOnly?: boolean;
    /**
     * Path within the volume from which the container's volume should be mounted. Defaults to "" (volume's root).
     */
    subPath?: string;
    /**
     * Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to "" (volume's root). SubPathExpr and SubPath are mutually exclusive.
     */
    subPathExpr?: string;
}
/**
 * VolumeNodeAffinity defines constraints that limit what nodes this volume can be accessed from.
 */
export type VolumeNodeAffinity = {
    /**
     * required specifies hard node constraints that must be met.
     */
    required?: NodeSelector;
}
/**
 * Projection that may be projected along with other supported volume types
 */
export type VolumeProjection = {
    /**
     * configMap information about the configMap data to project
     */
    configMap?: ConfigMapProjection;
    /**
     * downwardAPI information about the downwardAPI data to project
     */
    downwardAPI?: DownwardAPIProjection;
    /**
     * secret information about the secret data to project
     */
    secret?: SecretProjection;
    /**
     * serviceAccountToken is information about the serviceAccountToken data to project
     */
    serviceAccountToken?: ServiceAccountTokenProjection;
}
/**
 * Represents a vSphere volume resource.
 */
export type VsphereVirtualDiskVolumeSource = {
    /**
     * fsType is filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
     */
    fsType?: string;
    /**
     * storagePolicyID is the storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.
     */
    storagePolicyID?: string;
    /**
     * storagePolicyName is the storage Policy Based Management (SPBM) profile name.
     */
    storagePolicyName?: string;
    /**
     * volumePath is the path that identifies vSphere volume vmdk
     */
    volumePath: string;
}
/**
 * The weights of all of the matched WeightedPodAffinityTerm fields are added per-node to find the most preferred node(s)
 */
export type WeightedPodAffinityTerm = {
    /**
     * Required. A pod affinity term, associated with the corresponding weight.
     */
    podAffinityTerm: PodAffinityTerm;
    /**
     * weight associated with matching the corresponding podAffinityTerm, in the range 1-100.
     */
    weight: number;
}
/**
 * WindowsSecurityContextOptions contain Windows-specific options and credentials.
 */
export type WindowsSecurityContextOptions = {
    /**
     * GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.
     */
    gmsaCredentialSpec?: string;
    /**
     * GMSACredentialSpecName is the name of the GMSA credential spec to use.
     */
    gmsaCredentialSpecName?: string;
    /**
     * HostProcess determines if a container should be run as a 'Host Process' container. This field is alpha-level and will only be honored by components that enable the WindowsHostProcessContainers feature flag. Setting this field without the feature flag will result in errors when validating the Pod. All of a Pod's containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).  In addition, if HostProcess is true then HostNetwork must also be set to true.
     */
    hostProcess?: boolean;
    /**
     * The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
     */
    runAsUserName?: string;
}
