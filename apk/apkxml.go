package apk

import (
	"github.com/BlackLee123/apkparser"
)

// Instrumentation is an application instrumentation code.
type Instrumentation struct {
	Name            apkparser.String `xml:"http://schemas.android.com/apk/res/android name,attr"`
	Target          apkparser.String `xml:"http://schemas.android.com/apk/res/android targetPackage,attr"`
	HandleProfiling apkparser.Bool   `xml:"http://schemas.android.com/apk/res/android handleProfiling,attr"`
	FunctionalTest  apkparser.Bool   `xml:"http://schemas.android.com/apk/res/android functionalTest,attr"`
}

// ActivityAction is an action of an activity.
type ActivityAction struct {
	Name apkparser.String `xml:"http://schemas.android.com/apk/res/android name,attr"`
}

// ActivityCategory is a category of an activity.
type ActivityCategory struct {
	Name apkparser.String `xml:"http://schemas.android.com/apk/res/android name,attr"`
}

// ActivityIntentFilter is an apkparser.Int32ent filter of an activity.
type ActivityIntentFilter struct {
	Actions    []ActivityAction   `xml:"action"`
	Categories []ActivityCategory `xml:"category"`
}

// AppActivity is an activity in an application.
type AppActivity struct {
	Theme             apkparser.String       `xml:"http://schemas.android.com/apk/res/android theme,attr"`
	Name              apkparser.String       `xml:"http://schemas.android.com/apk/res/android name,attr"`
	Label             apkparser.String       `xml:"http://schemas.android.com/apk/res/android label,attr"`
	ScreenOrientation apkparser.String       `xml:"http://schemas.android.com/apk/res/android screenOrientation,attr"`
	IntentFilters     []ActivityIntentFilter `xml:"intent-filter"`
}

// AppActivityAlias https://developer.android.com/guide/topics/manifest/activity-alias-element
type AppActivityAlias struct {
	Name           apkparser.String       `xml:"http://schemas.android.com/apk/res/android name,attr"`
	Label          apkparser.String       `xml:"http://schemas.android.com/apk/res/android label,attr"`
	TargetActivity apkparser.String       `xml:"http://schemas.android.com/apk/res/android targetActivity,attr"`
	IntentFilters  []ActivityIntentFilter `xml:"intent-filter"`
}

// MetaData is a metadata in an application.
type MetaData struct {
	Name  apkparser.String `xml:"http://schemas.android.com/apk/res/android name,attr"`
	Value apkparser.String `xml:"http://schemas.android.com/apk/res/android value,attr"`
}

// Application is an application in an APK.
type Application struct {
	AllowTaskReparenting  apkparser.Bool     `xml:"http://schemas.android.com/apk/res/android allowTaskReparenting,attr"`
	AllowBackup           apkparser.Bool     `xml:"http://schemas.android.com/apk/res/android allowBackup,attr"`
	BackupAgent           apkparser.String   `xml:"http://schemas.android.com/apk/res/android backupAgent,attr"`
	Debuggable            apkparser.Bool     `xml:"http://schemas.android.com/apk/res/android debuggable,attr"`
	Description           apkparser.String   `xml:"http://schemas.android.com/apk/res/android description,attr"`
	Enabled               apkparser.Bool     `xml:"http://schemas.android.com/apk/res/android enabled,attr"`
	HasCode               apkparser.Bool     `xml:"http://schemas.android.com/apk/res/android hasCode,attr"`
	HardwareAccelerated   apkparser.Bool     `xml:"http://schemas.android.com/apk/res/android hardwareAccelerated,attr"`
	Icon                  apkparser.String   `xml:"http://schemas.android.com/apk/res/android icon,attr"`
	KillAfterRestore      apkparser.Bool     `xml:"http://schemas.android.com/apk/res/android killAfterRestore,attr"`
	LargeHeap             apkparser.Bool     `xml:"http://schemas.android.com/apk/res/android largeHeap,attr"`
	Label                 apkparser.String   `xml:"http://schemas.android.com/apk/res/android label,attr"`
	Logo                  apkparser.String   `xml:"http://schemas.android.com/apk/res/android logo,attr"`
	ManageSpaceActivity   apkparser.String   `xml:"http://schemas.android.com/apk/res/android manageSpaceActivity,attr"`
	Name                  apkparser.String   `xml:"http://schemas.android.com/apk/res/android name,attr"`
	Permission            apkparser.String   `xml:"http://schemas.android.com/apk/res/android permission,attr"`
	Persistent            apkparser.Bool     `xml:"http://schemas.android.com/apk/res/android persistent,attr"`
	Process               apkparser.String   `xml:"http://schemas.android.com/apk/res/android process,attr"`
	RestoreAnyVersion     apkparser.Bool     `xml:"http://schemas.android.com/apk/res/android restoreAnyVersion,attr"`
	RequiredAccountType   apkparser.String   `xml:"http://schemas.android.com/apk/res/android requiredAccountType,attr"`
	RestrictedAccountType apkparser.String   `xml:"http://schemas.android.com/apk/res/android restrictedAccountType,attr"`
	SupportsRtl           apkparser.Bool     `xml:"http://schemas.android.com/apk/res/android supportsRtl,attr"`
	TaskAffinity          apkparser.String   `xml:"http://schemas.android.com/apk/res/android taskAffinity,attr"`
	TestOnly              apkparser.Bool     `xml:"http://schemas.android.com/apk/res/android testOnly,attr"`
	Theme                 apkparser.String   `xml:"http://schemas.android.com/apk/res/android theme,attr"`
	UIOptions             apkparser.String   `xml:"http://schemas.android.com/apk/res/android uiOptions,attr"`
	VMSafeMode            apkparser.Bool     `xml:"http://schemas.android.com/apk/res/android vmSafeMode,attr"`
	Activities            []AppActivity      `xml:"activity"`
	ActivityAliases       []AppActivityAlias `xml:"activity-alias"`
	MetaData              []MetaData         `xml:"meta-data"`
}

// UsesSDK is target SDK version.
type UsesSDK struct {
	Min    apkparser.Int32 `xml:"http://schemas.android.com/apk/res/android minSdkVersion,attr"`
	Target apkparser.Int32 `xml:"http://schemas.android.com/apk/res/android targetSdkVersion,attr"`
	Max    apkparser.Int32 `xml:"http://schemas.android.com/apk/res/android maxSdkVersion,attr"`
}

// UsesPermission is user grant the system permission.
type UsesPermission struct {
	Name apkparser.String `xml:"http://schemas.android.com/apk/res/android name,attr"`
	Max  apkparser.Int32  `xml:"http://schemas.android.com/apk/res/android maxSdkVersion,attr"`
}

// Manifest is a manifest of an APK.
type Manifest struct {
	Package         apkparser.String `xml:"package,attr"`
	VersionCode     apkparser.Int32  `xml:"http://schemas.android.com/apk/res/android versionCode,attr"`
	VersionName     apkparser.String `xml:"http://schemas.android.com/apk/res/android versionName,attr"`
	App             Application      `xml:"application"`
	Instrument      Instrumentation  `xml:"instrumentation"`
	SDK             UsesSDK          `xml:"uses-sdk"`
	UsesPermissions []UsesPermission `xml:"uses-permission"`
}
