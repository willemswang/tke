// +build !ignore_autogenerated

/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2020 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by conversion-gen. DO NOT EDIT.

package v1

import (
	unsafe "unsafe"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	application "tkestack.io/tke/api/application"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*App)(nil), (*application.App)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_App_To_application_App(a.(*App), b.(*application.App), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*application.App)(nil), (*App)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_application_App_To_v1_App(a.(*application.App), b.(*App), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AppHistory)(nil), (*application.AppHistory)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_AppHistory_To_application_AppHistory(a.(*AppHistory), b.(*application.AppHistory), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*application.AppHistory)(nil), (*AppHistory)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_application_AppHistory_To_v1_AppHistory(a.(*application.AppHistory), b.(*AppHistory), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AppHistorySpec)(nil), (*application.AppHistorySpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_AppHistorySpec_To_application_AppHistorySpec(a.(*AppHistorySpec), b.(*application.AppHistorySpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*application.AppHistorySpec)(nil), (*AppHistorySpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_application_AppHistorySpec_To_v1_AppHistorySpec(a.(*application.AppHistorySpec), b.(*AppHistorySpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AppList)(nil), (*application.AppList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_AppList_To_application_AppList(a.(*AppList), b.(*application.AppList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*application.AppList)(nil), (*AppList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_application_AppList_To_v1_AppList(a.(*application.AppList), b.(*AppList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AppResource)(nil), (*application.AppResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_AppResource_To_application_AppResource(a.(*AppResource), b.(*application.AppResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*application.AppResource)(nil), (*AppResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_application_AppResource_To_v1_AppResource(a.(*application.AppResource), b.(*AppResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AppResourceSpec)(nil), (*application.AppResourceSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_AppResourceSpec_To_application_AppResourceSpec(a.(*AppResourceSpec), b.(*application.AppResourceSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*application.AppResourceSpec)(nil), (*AppResourceSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_application_AppResourceSpec_To_v1_AppResourceSpec(a.(*application.AppResourceSpec), b.(*AppResourceSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AppSpec)(nil), (*application.AppSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_AppSpec_To_application_AppSpec(a.(*AppSpec), b.(*application.AppSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*application.AppSpec)(nil), (*AppSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_application_AppSpec_To_v1_AppSpec(a.(*application.AppSpec), b.(*AppSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AppStatus)(nil), (*application.AppStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_AppStatus_To_application_AppStatus(a.(*AppStatus), b.(*application.AppStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*application.AppStatus)(nil), (*AppStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_application_AppStatus_To_v1_AppStatus(a.(*application.AppStatus), b.(*AppStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AppValues)(nil), (*application.AppValues)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_AppValues_To_application_AppValues(a.(*AppValues), b.(*application.AppValues), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*application.AppValues)(nil), (*AppValues)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_application_AppValues_To_v1_AppValues(a.(*application.AppValues), b.(*AppValues), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Chart)(nil), (*application.Chart)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Chart_To_application_Chart(a.(*Chart), b.(*application.Chart), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*application.Chart)(nil), (*Chart)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_application_Chart_To_v1_Chart(a.(*application.Chart), b.(*Chart), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ConfigMap)(nil), (*application.ConfigMap)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ConfigMap_To_application_ConfigMap(a.(*ConfigMap), b.(*application.ConfigMap), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*application.ConfigMap)(nil), (*ConfigMap)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_application_ConfigMap_To_v1_ConfigMap(a.(*application.ConfigMap), b.(*ConfigMap), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ConfigMapList)(nil), (*application.ConfigMapList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ConfigMapList_To_application_ConfigMapList(a.(*ConfigMapList), b.(*application.ConfigMapList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*application.ConfigMapList)(nil), (*ConfigMapList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_application_ConfigMapList_To_v1_ConfigMapList(a.(*application.ConfigMapList), b.(*ConfigMapList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*History)(nil), (*application.History)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_History_To_application_History(a.(*History), b.(*application.History), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*application.History)(nil), (*History)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_application_History_To_v1_History(a.(*application.History), b.(*History), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RollbackProxyOptions)(nil), (*application.RollbackProxyOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RollbackProxyOptions_To_application_RollbackProxyOptions(a.(*RollbackProxyOptions), b.(*application.RollbackProxyOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*application.RollbackProxyOptions)(nil), (*RollbackProxyOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_application_RollbackProxyOptions_To_v1_RollbackProxyOptions(a.(*application.RollbackProxyOptions), b.(*RollbackProxyOptions), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1_App_To_application_App(in *App, out *application.App, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_AppSpec_To_application_AppSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_AppStatus_To_application_AppStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_App_To_application_App is an autogenerated conversion function.
func Convert_v1_App_To_application_App(in *App, out *application.App, s conversion.Scope) error {
	return autoConvert_v1_App_To_application_App(in, out, s)
}

func autoConvert_application_App_To_v1_App(in *application.App, out *App, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_application_AppSpec_To_v1_AppSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_application_AppStatus_To_v1_AppStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_application_App_To_v1_App is an autogenerated conversion function.
func Convert_application_App_To_v1_App(in *application.App, out *App, s conversion.Scope) error {
	return autoConvert_application_App_To_v1_App(in, out, s)
}

func autoConvert_v1_AppHistory_To_application_AppHistory(in *AppHistory, out *application.AppHistory, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_AppHistorySpec_To_application_AppHistorySpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_AppHistory_To_application_AppHistory is an autogenerated conversion function.
func Convert_v1_AppHistory_To_application_AppHistory(in *AppHistory, out *application.AppHistory, s conversion.Scope) error {
	return autoConvert_v1_AppHistory_To_application_AppHistory(in, out, s)
}

func autoConvert_application_AppHistory_To_v1_AppHistory(in *application.AppHistory, out *AppHistory, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_application_AppHistorySpec_To_v1_AppHistorySpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_application_AppHistory_To_v1_AppHistory is an autogenerated conversion function.
func Convert_application_AppHistory_To_v1_AppHistory(in *application.AppHistory, out *AppHistory, s conversion.Scope) error {
	return autoConvert_application_AppHistory_To_v1_AppHistory(in, out, s)
}

func autoConvert_v1_AppHistorySpec_To_application_AppHistorySpec(in *AppHistorySpec, out *application.AppHistorySpec, s conversion.Scope) error {
	out.Type = application.AppType(in.Type)
	out.TenantID = in.TenantID
	out.Name = in.Name
	out.TargetCluster = in.TargetCluster
	out.Histories = *(*[]application.History)(unsafe.Pointer(&in.Histories))
	return nil
}

// Convert_v1_AppHistorySpec_To_application_AppHistorySpec is an autogenerated conversion function.
func Convert_v1_AppHistorySpec_To_application_AppHistorySpec(in *AppHistorySpec, out *application.AppHistorySpec, s conversion.Scope) error {
	return autoConvert_v1_AppHistorySpec_To_application_AppHistorySpec(in, out, s)
}

func autoConvert_application_AppHistorySpec_To_v1_AppHistorySpec(in *application.AppHistorySpec, out *AppHistorySpec, s conversion.Scope) error {
	out.Type = AppType(in.Type)
	out.TenantID = in.TenantID
	out.Name = in.Name
	out.TargetCluster = in.TargetCluster
	out.Histories = *(*[]History)(unsafe.Pointer(&in.Histories))
	return nil
}

// Convert_application_AppHistorySpec_To_v1_AppHistorySpec is an autogenerated conversion function.
func Convert_application_AppHistorySpec_To_v1_AppHistorySpec(in *application.AppHistorySpec, out *AppHistorySpec, s conversion.Scope) error {
	return autoConvert_application_AppHistorySpec_To_v1_AppHistorySpec(in, out, s)
}

func autoConvert_v1_AppList_To_application_AppList(in *AppList, out *application.AppList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]application.App)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_AppList_To_application_AppList is an autogenerated conversion function.
func Convert_v1_AppList_To_application_AppList(in *AppList, out *application.AppList, s conversion.Scope) error {
	return autoConvert_v1_AppList_To_application_AppList(in, out, s)
}

func autoConvert_application_AppList_To_v1_AppList(in *application.AppList, out *AppList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]App)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_application_AppList_To_v1_AppList is an autogenerated conversion function.
func Convert_application_AppList_To_v1_AppList(in *application.AppList, out *AppList, s conversion.Scope) error {
	return autoConvert_application_AppList_To_v1_AppList(in, out, s)
}

func autoConvert_v1_AppResource_To_application_AppResource(in *AppResource, out *application.AppResource, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_AppResourceSpec_To_application_AppResourceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_AppResource_To_application_AppResource is an autogenerated conversion function.
func Convert_v1_AppResource_To_application_AppResource(in *AppResource, out *application.AppResource, s conversion.Scope) error {
	return autoConvert_v1_AppResource_To_application_AppResource(in, out, s)
}

func autoConvert_application_AppResource_To_v1_AppResource(in *application.AppResource, out *AppResource, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_application_AppResourceSpec_To_v1_AppResourceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_application_AppResource_To_v1_AppResource is an autogenerated conversion function.
func Convert_application_AppResource_To_v1_AppResource(in *application.AppResource, out *AppResource, s conversion.Scope) error {
	return autoConvert_application_AppResource_To_v1_AppResource(in, out, s)
}

func autoConvert_v1_AppResourceSpec_To_application_AppResourceSpec(in *AppResourceSpec, out *application.AppResourceSpec, s conversion.Scope) error {
	out.Type = application.AppType(in.Type)
	out.TenantID = in.TenantID
	out.Name = in.Name
	out.TargetCluster = in.TargetCluster
	out.Resources = *(*application.Resources)(unsafe.Pointer(&in.Resources))
	return nil
}

// Convert_v1_AppResourceSpec_To_application_AppResourceSpec is an autogenerated conversion function.
func Convert_v1_AppResourceSpec_To_application_AppResourceSpec(in *AppResourceSpec, out *application.AppResourceSpec, s conversion.Scope) error {
	return autoConvert_v1_AppResourceSpec_To_application_AppResourceSpec(in, out, s)
}

func autoConvert_application_AppResourceSpec_To_v1_AppResourceSpec(in *application.AppResourceSpec, out *AppResourceSpec, s conversion.Scope) error {
	out.Type = AppType(in.Type)
	out.TenantID = in.TenantID
	out.Name = in.Name
	out.TargetCluster = in.TargetCluster
	out.Resources = *(*Resources)(unsafe.Pointer(&in.Resources))
	return nil
}

// Convert_application_AppResourceSpec_To_v1_AppResourceSpec is an autogenerated conversion function.
func Convert_application_AppResourceSpec_To_v1_AppResourceSpec(in *application.AppResourceSpec, out *AppResourceSpec, s conversion.Scope) error {
	return autoConvert_application_AppResourceSpec_To_v1_AppResourceSpec(in, out, s)
}

func autoConvert_v1_AppSpec_To_application_AppSpec(in *AppSpec, out *application.AppSpec, s conversion.Scope) error {
	out.Type = application.AppType(in.Type)
	out.TenantID = in.TenantID
	out.Name = in.Name
	out.TargetCluster = in.TargetCluster
	if err := Convert_v1_Chart_To_application_Chart(&in.Chart, &out.Chart, s); err != nil {
		return err
	}
	if err := Convert_v1_AppValues_To_application_AppValues(&in.Values, &out.Values, s); err != nil {
		return err
	}
	out.Finalizers = *(*[]application.FinalizerName)(unsafe.Pointer(&in.Finalizers))
	return nil
}

// Convert_v1_AppSpec_To_application_AppSpec is an autogenerated conversion function.
func Convert_v1_AppSpec_To_application_AppSpec(in *AppSpec, out *application.AppSpec, s conversion.Scope) error {
	return autoConvert_v1_AppSpec_To_application_AppSpec(in, out, s)
}

func autoConvert_application_AppSpec_To_v1_AppSpec(in *application.AppSpec, out *AppSpec, s conversion.Scope) error {
	out.Type = AppType(in.Type)
	out.TenantID = in.TenantID
	out.Name = in.Name
	out.TargetCluster = in.TargetCluster
	if err := Convert_application_Chart_To_v1_Chart(&in.Chart, &out.Chart, s); err != nil {
		return err
	}
	if err := Convert_application_AppValues_To_v1_AppValues(&in.Values, &out.Values, s); err != nil {
		return err
	}
	out.Finalizers = *(*[]FinalizerName)(unsafe.Pointer(&in.Finalizers))
	return nil
}

// Convert_application_AppSpec_To_v1_AppSpec is an autogenerated conversion function.
func Convert_application_AppSpec_To_v1_AppSpec(in *application.AppSpec, out *AppSpec, s conversion.Scope) error {
	return autoConvert_application_AppSpec_To_v1_AppSpec(in, out, s)
}

func autoConvert_v1_AppStatus_To_application_AppStatus(in *AppStatus, out *application.AppStatus, s conversion.Scope) error {
	out.Phase = application.AppPhase(in.Phase)
	out.ObservedGeneration = in.ObservedGeneration
	out.ReleaseStatus = in.ReleaseStatus
	out.ReleaseLastUpdated = in.ReleaseLastUpdated
	out.Revision = in.Revision
	out.RollbackRevision = in.RollbackRevision
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_v1_AppStatus_To_application_AppStatus is an autogenerated conversion function.
func Convert_v1_AppStatus_To_application_AppStatus(in *AppStatus, out *application.AppStatus, s conversion.Scope) error {
	return autoConvert_v1_AppStatus_To_application_AppStatus(in, out, s)
}

func autoConvert_application_AppStatus_To_v1_AppStatus(in *application.AppStatus, out *AppStatus, s conversion.Scope) error {
	out.Phase = AppPhase(in.Phase)
	out.ObservedGeneration = in.ObservedGeneration
	out.ReleaseStatus = in.ReleaseStatus
	out.ReleaseLastUpdated = in.ReleaseLastUpdated
	out.Revision = in.Revision
	out.RollbackRevision = in.RollbackRevision
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_application_AppStatus_To_v1_AppStatus is an autogenerated conversion function.
func Convert_application_AppStatus_To_v1_AppStatus(in *application.AppStatus, out *AppStatus, s conversion.Scope) error {
	return autoConvert_application_AppStatus_To_v1_AppStatus(in, out, s)
}

func autoConvert_v1_AppValues_To_application_AppValues(in *AppValues, out *application.AppValues, s conversion.Scope) error {
	out.RawValuesType = application.RawValuesType(in.RawValuesType)
	out.RawValues = in.RawValues
	out.Values = *(*[]string)(unsafe.Pointer(&in.Values))
	return nil
}

// Convert_v1_AppValues_To_application_AppValues is an autogenerated conversion function.
func Convert_v1_AppValues_To_application_AppValues(in *AppValues, out *application.AppValues, s conversion.Scope) error {
	return autoConvert_v1_AppValues_To_application_AppValues(in, out, s)
}

func autoConvert_application_AppValues_To_v1_AppValues(in *application.AppValues, out *AppValues, s conversion.Scope) error {
	out.RawValuesType = RawValuesType(in.RawValuesType)
	out.RawValues = in.RawValues
	out.Values = *(*[]string)(unsafe.Pointer(&in.Values))
	return nil
}

// Convert_application_AppValues_To_v1_AppValues is an autogenerated conversion function.
func Convert_application_AppValues_To_v1_AppValues(in *application.AppValues, out *AppValues, s conversion.Scope) error {
	return autoConvert_application_AppValues_To_v1_AppValues(in, out, s)
}

func autoConvert_v1_Chart_To_application_Chart(in *Chart, out *application.Chart, s conversion.Scope) error {
	out.TenantID = in.TenantID
	out.ChartGroupName = in.ChartGroupName
	out.ChartName = in.ChartName
	out.ChartVersion = in.ChartVersion
	return nil
}

// Convert_v1_Chart_To_application_Chart is an autogenerated conversion function.
func Convert_v1_Chart_To_application_Chart(in *Chart, out *application.Chart, s conversion.Scope) error {
	return autoConvert_v1_Chart_To_application_Chart(in, out, s)
}

func autoConvert_application_Chart_To_v1_Chart(in *application.Chart, out *Chart, s conversion.Scope) error {
	out.TenantID = in.TenantID
	out.ChartGroupName = in.ChartGroupName
	out.ChartName = in.ChartName
	out.ChartVersion = in.ChartVersion
	return nil
}

// Convert_application_Chart_To_v1_Chart is an autogenerated conversion function.
func Convert_application_Chart_To_v1_Chart(in *application.Chart, out *Chart, s conversion.Scope) error {
	return autoConvert_application_Chart_To_v1_Chart(in, out, s)
}

func autoConvert_v1_ConfigMap_To_application_ConfigMap(in *ConfigMap, out *application.ConfigMap, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Data = *(*map[string]string)(unsafe.Pointer(&in.Data))
	out.BinaryData = *(*map[string][]byte)(unsafe.Pointer(&in.BinaryData))
	return nil
}

// Convert_v1_ConfigMap_To_application_ConfigMap is an autogenerated conversion function.
func Convert_v1_ConfigMap_To_application_ConfigMap(in *ConfigMap, out *application.ConfigMap, s conversion.Scope) error {
	return autoConvert_v1_ConfigMap_To_application_ConfigMap(in, out, s)
}

func autoConvert_application_ConfigMap_To_v1_ConfigMap(in *application.ConfigMap, out *ConfigMap, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Data = *(*map[string]string)(unsafe.Pointer(&in.Data))
	out.BinaryData = *(*map[string][]byte)(unsafe.Pointer(&in.BinaryData))
	return nil
}

// Convert_application_ConfigMap_To_v1_ConfigMap is an autogenerated conversion function.
func Convert_application_ConfigMap_To_v1_ConfigMap(in *application.ConfigMap, out *ConfigMap, s conversion.Scope) error {
	return autoConvert_application_ConfigMap_To_v1_ConfigMap(in, out, s)
}

func autoConvert_v1_ConfigMapList_To_application_ConfigMapList(in *ConfigMapList, out *application.ConfigMapList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]application.ConfigMap)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_ConfigMapList_To_application_ConfigMapList is an autogenerated conversion function.
func Convert_v1_ConfigMapList_To_application_ConfigMapList(in *ConfigMapList, out *application.ConfigMapList, s conversion.Scope) error {
	return autoConvert_v1_ConfigMapList_To_application_ConfigMapList(in, out, s)
}

func autoConvert_application_ConfigMapList_To_v1_ConfigMapList(in *application.ConfigMapList, out *ConfigMapList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]ConfigMap)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_application_ConfigMapList_To_v1_ConfigMapList is an autogenerated conversion function.
func Convert_application_ConfigMapList_To_v1_ConfigMapList(in *application.ConfigMapList, out *ConfigMapList, s conversion.Scope) error {
	return autoConvert_application_ConfigMapList_To_v1_ConfigMapList(in, out, s)
}

func autoConvert_v1_History_To_application_History(in *History, out *application.History, s conversion.Scope) error {
	out.Revision = in.Revision
	out.Updated = in.Updated
	out.Status = in.Status
	out.Chart = in.Chart
	out.AppVersion = in.AppVersion
	out.Description = in.Description
	return nil
}

// Convert_v1_History_To_application_History is an autogenerated conversion function.
func Convert_v1_History_To_application_History(in *History, out *application.History, s conversion.Scope) error {
	return autoConvert_v1_History_To_application_History(in, out, s)
}

func autoConvert_application_History_To_v1_History(in *application.History, out *History, s conversion.Scope) error {
	out.Revision = in.Revision
	out.Updated = in.Updated
	out.Status = in.Status
	out.Chart = in.Chart
	out.AppVersion = in.AppVersion
	out.Description = in.Description
	return nil
}

// Convert_application_History_To_v1_History is an autogenerated conversion function.
func Convert_application_History_To_v1_History(in *application.History, out *History, s conversion.Scope) error {
	return autoConvert_application_History_To_v1_History(in, out, s)
}

func autoConvert_v1_RollbackProxyOptions_To_application_RollbackProxyOptions(in *RollbackProxyOptions, out *application.RollbackProxyOptions, s conversion.Scope) error {
	out.Revision = in.Revision
	out.Cluster = in.Cluster
	return nil
}

// Convert_v1_RollbackProxyOptions_To_application_RollbackProxyOptions is an autogenerated conversion function.
func Convert_v1_RollbackProxyOptions_To_application_RollbackProxyOptions(in *RollbackProxyOptions, out *application.RollbackProxyOptions, s conversion.Scope) error {
	return autoConvert_v1_RollbackProxyOptions_To_application_RollbackProxyOptions(in, out, s)
}

func autoConvert_application_RollbackProxyOptions_To_v1_RollbackProxyOptions(in *application.RollbackProxyOptions, out *RollbackProxyOptions, s conversion.Scope) error {
	out.Revision = in.Revision
	out.Cluster = in.Cluster
	return nil
}

// Convert_application_RollbackProxyOptions_To_v1_RollbackProxyOptions is an autogenerated conversion function.
func Convert_application_RollbackProxyOptions_To_v1_RollbackProxyOptions(in *application.RollbackProxyOptions, out *RollbackProxyOptions, s conversion.Scope) error {
	return autoConvert_application_RollbackProxyOptions_To_v1_RollbackProxyOptions(in, out, s)
}