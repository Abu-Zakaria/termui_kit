package main

func NewKit() *Kit {
	return &Kit{
		Elements: []Element{},
	}
}

func (kit *Kit) AddElement(element Element) {
	kit.Elements = append(kit.Elements, element)
}
