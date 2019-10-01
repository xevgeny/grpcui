package standalone

const IndexTemplateName = "index-template.html"
const ListTemplateName = "list-template.html"

func IndexTemplate() []byte {
	return MustAsset(IndexTemplateName)
}

func ListTemplate() []byte {
	return MustAsset(ListTemplateName)
}
