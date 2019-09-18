package v1alpha1


var SchemeGroupVersion = schema.GroupVersion{
  Group: "bar.com",
  Version: "v1"
}

var (
  SchemeBuilder runtime.SchemeBuilder
  localSchemeBuilder = &SchemeBuilder
  AddToScheme = localSchemeBuilder.AddToScheme
)

func init()  {
  localSchemeBuilder.Regisrer(addKnownTypes)

}

func Resource(resource string) schema.GroupResource  {
  return SchemeGroupVersion.WithResource(resource).GroupResource()

}

func addKnownTypes(scheme *runtime.Scheme) error  {
  scheme.AddKnownTypes(
    SchemeGroupVersion,
    &HelloType{},
    &HelloTypeList{}
  )

  scheme.AddKnownTypes(
    SchemeGroupVersion,
    &metav1.Status{}
  )

  metav1.AddToGroupVersion(
    scheme,
    SchemeGroupVersion
  )
  return nil

}
