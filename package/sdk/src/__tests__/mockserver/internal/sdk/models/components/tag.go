// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Tag struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

func (o *Tag) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Tag) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
