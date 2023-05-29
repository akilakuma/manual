package mysql

func (o *Orm) Association(column string) *Association {
	association := o.db.Association(column)

	a := &Association{
		Association: association,
	}

	return a
}

func (a *Association) Append(values ...interface{}) *Association {
	return &Association{Association: a.Association.Append(values...)}
}

func (a *Association) Clear() *Association {
	a.Association = a.Association.Clear()

	return a
}

func (a *Association) Count() int {
	return a.Association.Count()
}

func (a *Association) Delete(values ...interface{}) *Association {
	return &Association{Association: a.Association.Delete(values...)}
}

func (a *Association) Find(value interface{}) *Association {
	return &Association{Association: a.Association.Find(value)}
}

func (a *Association) Replace(values ...interface{}) *Association {
	return &Association{Association: a.Association.Replace(values...)}
}

func (a *Association) Error() error {
	return a.Association.Error
}
