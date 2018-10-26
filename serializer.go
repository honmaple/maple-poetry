/*********************************************************************************
 Copyright © 2018 jianglin
 File Name: serializer.go
 Author: jianglin
 Email: mail@honmaple.com
 Created: 2018-09-12 15:04:57 (CST)
 Last Update: Monday 2018-10-22 11:52:08 (CST)
		  By:
 Description:
 *********************************************************************************/
package main

// PoemSerializer ..
type PoemSerializer struct {
	Instances []*Poem
	Instance  *Poem
}

// Data ..
func (s *PoemSerializer) Data() []map[string]interface{} {
	ins := make([]map[string]interface{}, 0)
	if s.Instances != nil {
		for _, instance := range s.Instances {
			ins = append(ins, instance.Serializer())
		}
		return ins
	}
	ins = append(ins, s.Instance.Serializer())
	return ins
}

// AuthorSerializer ..
type AuthorSerializer struct {
	Instances []*Author
	Instance  *Author
}

// Data ..
func (s *AuthorSerializer) Data() []map[string]interface{} {
	ins := make([]map[string]interface{}, 0)
	if s.Instances != nil {
		for _, instance := range s.Instances {
			ins = append(ins, instance.Serializer())
		}
		return ins
	}
	ins = append(ins, s.Instance.Serializer())
	return ins
}
