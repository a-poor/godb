package godb

type DataType int

const (
	TypeString DataType = iota
	TypeInt
	TypeFloat
	TypeBool
	TypeCategory
)

type Row struct {
	F string `godb:"nullable"`
}

type Column struct {
	Index      int
	Name       string
	Type       DataType
	IsNullable bool
}

func (c Column) ToBinary() []byte {
	return nil
}

type DB struct {
	Filename   string
	Columns    []Column
	Categories map[string]map[uint8]string
}

// func (db *DB) createHeader() []byte {
// 	return make([]byte, 0)
// }

// func (db *DB) loadHeader() error {
// 	return nil
// }

/*
 *
 *
 */
func NewDB(filename string, rowType interface{}) (*DB, error) {
	return &DB{Filename: filename}, nil
}

// func getColsFromStruct(t interface{}) ([]Column, error) {
// 	return nil, nil
// }

// func main() {
// 	r := Row{}
// 	rt := reflect.TypeOf(r)
// 	for i := 0; i < rt.NumField(); i++ {
// 		f := rt.Field(i)
// 		fmt.Println(f.Name)
// 		for k, v := range f.Tag {
// 			fmt.Printf("    %s = %s\n", string(k), v)
// 		}
// 	}
// }
