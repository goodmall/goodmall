package base

// TODO  可以作为基础类 供所有的模块使用 考虑提取到 base/usercase/query.go 文件去

/**
* 拷贝自： https://github.com/CaptainCodeman/clean-go/blob/master/engine/query.go

  该类表示的是应用级的查询规范  在存储层需要翻译为特定存储的criterion

书籍参考：<<Domain Driven Design in PHP>>
  >	pass a criterion and let the Repository handle the implementation
details to successfully perform the operation. It might translate the criterion to SQL or ORM
queries or iterate over an in-memory collection.

  引入这个类的主要原因是 Repository跟DAO查询方法的重要区别  由于传统DAO 是面向table的 其查询方法比较
  发散 充斥着 findByXxx  根据查询场景会越来越多的 这不利于测试  在Repository中 为了消除这种情况查询接口
做了归一化  IRepo::Query(specification);

  特定的仓储实现 需要翻译为自己的查询条件 比如rdbms就是 sql片段    mongodb就是查询文档  纯内存实现就是对象规格
（Specification pattern   规格是一种类似过滤规范的接口 是一种简单的断言 接受一个领域对象并返回一个布尔值 在遍历集合
时会计算这个接口的特定函数的 只会返回true值的对象  类似filter函数传递了一个过滤器函数）
*/

// Direction represents a query sort direction
type Direction byte

const (
	// Ascending means going up, A-Z
	Ascending Direction = 1 << iota

	// Descending means reverse order, Z-A
	Descending
)

// Condition represents a filter comparison operation
// between a field and a value
type Condition byte

const (
	// Equal if it should be the same
	Equal Condition = 1 << iota

	// LessThan if it should be smaller
	LessThan

	// LessThanOrEqual if it should be smaller or equal
	LessThanOrEqual

	// GreaterThan if it should be larger
	GreaterThan

	// GreaterThanOrEqual if it should be equal or greater than
	GreaterThanOrEqual
)

type (
	// Query represents a query specification for filtering
	// sorting, paging and limiting the data requested
	Query struct {
		Name    string
		Offset  int
		Limit   int
		Filters []*Filter
		Orders  []*Order
	}

	// QueryBuilder helps with query creation
	QueryBuilder interface {
		Filter(property string, value interface{}) QueryBuilder
		Order(property string, direction Direction)
	}

	// Filter represents a filter operation on a single field
	Filter struct {
		Property  string
		Condition Condition
		Value     interface{}
	}

	// Order represents a sort operation on a single field
	Order struct {
		Property  string
		Direction Direction
	}
)

// NewQuery creates a new database query spec. The name is what
// the storage system should use to identify the types, usually
// a table or collection name.
func NewQuery(name string) *Query {
	return &Query{
		Name: name,
	}
}

// Filter adds a filter to the query
func (q *Query) Filter(property string, condition Condition, value interface{}) *Query {
	filter := NewFilter(property, condition, value)
	q.Filters = append(q.Filters, filter)
	return q
}

// Order adds a sort order to the query
func (q *Query) Order(property string, direction Direction) *Query {
	order := NewOrder(property, direction)
	q.Orders = append(q.Orders, order)
	return q
}

// Slice adds a slice operation to the query
func (q *Query) Slice(offset, limit int) *Query {
	q.Offset = offset
	q.Limit = limit
	return q
}

// NewFilter creates a new property filter
func NewFilter(property string, condition Condition, value interface{}) *Filter {
	return &Filter{
		Property:  property,
		Condition: condition,
		Value:     value,
	}
}

// NewOrder creates a new query order
func NewOrder(property string, direction Direction) *Order {
	return &Order{
		Property:  property,
		Direction: direction,
	}
}
