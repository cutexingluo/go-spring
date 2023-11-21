package parse

import (
	"github.com/cutexingluo/go-spring/common/base"
	"github.com/cutexingluo/go-spring/common/structure"
)

type DependencyChain struct {
	src string // source struct
	to  string // target struct , in beans
}

type DependencyNode struct {
	src  string // source struct
	to   string // target struct, in beans
	next int    // next node index
}

func NewDependencyNode(chain *DependencyChain, next int) DependencyNode {
	return DependencyNode{
		src:  chain.src,
		to:   chain.to,
		next: next,
	}
}

type BeanTopo struct {
	lineMap map[DependencyChain]bool //start
	total   int                      // line  total

	lines []DependencyNode    //from index 1 to total ,inclusive
	head  map[string]int      // head node index
	in    map[string]int      // in num , node count
	queue *structure.LinkList // the linked list
	//sorted []DependencyNode // sorted
}

// GetLineMap -  you can see the dependency line from lineMap
func (_this *BeanTopo) GetLineMap() map[DependencyChain]bool {
	return _this.lineMap
}

func (_this *DependencyChain) Equals(other DependencyChain) bool {
	return _this.src == other.src && _this.to == other.to
}

func NewBeanTopo() *BeanTopo {
	return &BeanTopo{
		lineMap: make(map[DependencyChain]bool),
		in:      make(map[string]int),
		queue:   structure.NewLinkList(),
		total:   0,
	}
}

// Add  1. 第一步添加
func (_this *BeanTopo) Add(src string, to string) {
	if src == "" || to == "" {
		return
	}
	_this.lineMap[DependencyChain{src: src, to: to}] = true
	_this.in[src] = 0 // init node count
	_this.in[to] = 0
}

func (_this *BeanTopo) addLine(chain *DependencyChain) {
	_this.total++
	_this.lines[_this.total] = NewDependencyNode(chain, _this.head[chain.src])
	_this.head[chain.src] = _this.total
	_this.in[chain.to]++ // 入度+1
}

// GetLines 2.获取所有依赖链
func (_this *BeanTopo) GetLines() []DependencyNode {
	if len(_this.lines) != 0 {
		return _this.lines
	}
	length := len(_this.lineMap)
	_this.lines = make([]DependencyNode, length+1)
	_this.head = make(map[string]int)
	for line := range _this.lineMap {
		_this.addLine(&line)
	}
	return _this.lines
}

// GetTopoBeans 3.获取 topo 序列, 优先度依次升高
func (_this *BeanTopo) GetTopoBeans() (sorted []string, err error) {
	if !_this.queue.IsEmpty() {
		_this.queue.Clear()
	}
	for s, cnt := range _this.in {
		if cnt == 0 {
			_this.queue.Offer(s)
		}
	}
	if _this.queue.IsEmpty() {
		if len(_this.in) == 0 {
			return nil, nil
		} else {
			return nil, &base.ErrIllegalState{ErrMsg: "Failed to initialize ALL beans,\n which may have formed circular dependencies"}
		}
	}
	cnt := 0
	n := len(_this.in)
	sorted = make([]string, n)
	for !_this.queue.IsEmpty() {
		now := _this.queue.Poll().(string)
		sorted[cnt] = now
		cnt++
		for i := _this.head[now]; i > 0; i = _this.lines[i].next {
			v := _this.lines[i].to
			_this.in[v]--
			if _this.in[v] == 0 {
				_this.queue.Offer(v)
			}
		}
	}
	if (cnt ^ n) != 0 {
		return nil, &base.ErrIllegalState{ErrMsg: " Failed to initialize beans,\n --> which may have formed circular dependencies"}
	}
	return sorted, nil
}

// Build After Adds, Build the topo from step 2 and step 3
func (_this *BeanTopo) Build() (sorted []string, err error) {
	_this.GetLines()
	return _this.GetTopoBeans()
}
