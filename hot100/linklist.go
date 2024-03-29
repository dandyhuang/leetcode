package leetcode_hot100

type ListNode struct {
	Val  int
	Next *ListNode
}

// 160. 相交链表
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	hA, hB := headA, headB
	for hA != hB {
		if hA == nil {
			hA = headB
		} else {
			hA = hA.Next
		}
		if hB == nil {
			hB = headA
		} else {
			hB = hB.Next
		}
	}
	return hA
}

// 206. 反转链表
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for cur := head; cur != nil; {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// 206. 反转链表 递归
func reverseListRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseListRecursion(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

// 234. 回文链表
func isPalindrome(head *ListNode) bool {
	var values []int
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}
	for i := 0; i < len(values)/2; i++ {
		if values[i] != values[len(values)-i-1] {
			return false
		}
	}
	return true
}

// 141. 环形链表
func hasCycle(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func isCycle(slow, fast *ListNode) (bool, *ListNode) {
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true, slow
		}
	}
	return false, nil
}

// 142. 环形链表 II
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	isC, node := isCycle(slow, fast)
	if isC {
		for head != nil {
			if head == node {
				return node
			}
			head = head.Next
			node = node.Next
		}
	}
	return nil
}

// 21. 合并两个有序链表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	if list1 != nil {
		cur.Next = list1
	}
	if list2 != nil {
		cur.Next = list2
	}
	return dummy.Next
}

// 2. 两数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	plus := 0
	for l1 != nil || l2 != nil {
		sum := plus
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		val := sum % 10
		plus = sum / 10
		cur.Next = &ListNode{Val: val}
		cur = cur.Next
	}
	if plus > 0 {
		cur.Next = &ListNode{Val: plus}
	}
	return dummy.Next
}

// 19. 删除链表的倒数第 N 个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy
	// size := 0
	// 这里还是少走了一位！！！
	//for fast != nil {
	//	fast = fast.Next
	//	size++
	//	if size == n {
	//		break
	//	}
	//}
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

// 24. 两两交换链表中的节点
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	for cur := dummy; cur.Next != nil && cur.Next.Next != nil; {
		q := cur.Next
		p := cur.Next.Next
		cur.Next = p
		cur = q
		q.Next = p.Next
		p.Next = q
	}
	return dummy.Next
}

// 24. 两两交换链表中的节点 递归
// 输入：head = [1,2,3,4]
// 输出：[2,1,4,3]
func swapPairsRecursion(head *ListNode) *ListNode {
	if head == nil && head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairsRecursion(next.Next)
	next.Next = head
	return next
}

// 25. K 个一组翻转链表
// 输入：head = [1,2,3,4,5], k = 2
// 输出：[2,1,4,3,5]
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	prevGroupTail := dummy
	for head != nil {
		groupHead := head
		groupTail := head
		for i := 1; i < k && groupTail != nil; i++ {
			groupTail = groupTail.Next
		}
		if groupTail == nil {
			break
		}
		nextGroupHead := groupTail.Next
		groupTail.Next = nil

		prevGroupTail.Next = reverseList(groupHead)
		groupHead.Next = nextGroupHead
		prevGroupTail = groupHead
		head = nextGroupHead
	}
	return dummy.Next
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 138. 随机链表的复制
func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node)
	for cur := head; cur != nil; cur = cur.Next {
		m[cur] = &Node{Val: cur.Val}
	}
	for cur := head; cur != nil; cur = cur.Next {
		if cur.Next != nil {
			m[cur].Next = m[cur.Next]
		}
		if cur.Random != nil {
			m[cur].Random = m[cur.Random]
		}
	}
	return m[head]
}

// 148. 排序链表
// 输入：head = [4,2,1,3]
// 输出：[1,2,3,4]
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 找到链表的中点
	pre := findMiddle(head)
	mid := pre.Next
	pre.Next = nil // 切断链表，分为两个独立的链表
	// 递归排序左右两部分链表
	left := sortList(head)
	right := sortList(mid)

	// 合并有序链表
	return mergeSort(left, right)
}

func findMiddle(head *ListNode) *ListNode {
	slow, fast := head, head
	var prev *ListNode

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	return prev
}

func mergeSort(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	if l1 != nil {
		current.Next = l1
	}

	if l2 != nil {
		current.Next = l2
	}

	return dummy.Next
}

func sortListV2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next // 偶数情况，需要返回中位数的前一个
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	left := sortList(head)
	right := sortList(mid)
	dump := &ListNode{}
	cur := dump
	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}
		cur = cur.Next
	}
	if left != nil {
		cur.Next = left
	}
	if right != nil {
		cur.Next = right
	}
	return dump.Next
}

// 23. 合并 K 个升序链表
func mergeTwoList(a *ListNode, b *ListNode) *ListNode {
	dump := &ListNode{}
	tmp := dump
	for a != nil && b != nil {
		if a.Val < b.Val {
			tmp.Next = a
			a = a.Next
		} else {
			tmp.Next = b
			b = b.Next
		}
		tmp = tmp.Next
	}
	if a != nil {
		tmp.Next = a
	}
	if b != nil {
		tmp.Next = b
	}
	return dump.Next
}

// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
// 输出：[1,1,2,3,4,4,5,6]
// 解释：链表数组如下：
// [
//
//	1->4->5,
//	1->3->4,
//	2->6
//
// ]
func mergeKLists(lists []*ListNode) *ListNode {
	var res *ListNode
	for i := range lists {
		res = mergeTwoList(res, lists[i])
	}
	return res
}

// NodeList 146. LRU 缓存
type NodeList struct {
	// 双向链表
	pre, next *NodeList
	k, v      int
}

type LRUCache struct {
	capacity, size int
	head, tail     *NodeList
	h              map[int]*NodeList
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{capacity: capacity}
	lru.head = &NodeList{}
	lru.tail = &NodeList{}
	lru.head.next = lru.tail
	lru.tail.pre = lru.head
	lru.h = make(map[int]*NodeList, 0)
	return lru
}

func (l *LRUCache) Put(key, value int) {
	if v, ok := l.h[key]; ok {
		v.v = value
		l.h[key] = v
		l.RemoveNode(v)
		l.AddNode(v)
	} else {
		if l.capacity == l.size {
			node := l.tail.pre
			delete(l.h, node.k)
			l.RemoveNode(node)
			l.size--
		}
		node := &NodeList{
			k: key,
			v: value,
		}
		l.AddNode(node)
		l.h[key] = node
		l.size++
	}
}

func (l *LRUCache) Get(key int) int {
	if v, ok := l.h[key]; ok {
		l.RemoveNode(v)
		l.AddNode(v)
		return v.v
	}
	return -1
}
func (l *LRUCache) AddNode(node *NodeList) {
	node.pre = l.head
	node.next = l.head.next
	l.head.next.pre = node
	l.head.next = node
}
func (l *LRUCache) RemoveNode(node *NodeList) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

/*
struct Node {
    // 双向链表
    Node* pre, *next;
    int key, val;
    Node():key(0), val(0),pre(nullptr), next(nullptr) {}
    Node(int key, int val):key(key), val(val),pre(nullptr), next(nullptr) {}
};
class LRUCache {
public:
    LRUCache(int capacity) {
        capacity_ = capacity;
        size_ = 0;
        head_ = new Node();
        tail_ = new Node();
        head_->next = tail_;
        tail_->pre = head_;
    }

    int get(int key) {
        if (h_.count(key)) {
            Node* node = h_[key];
            RemoveNode(node);
            AddNode(node);
            return node->val;
        }
        return -1;
    }

    void put(int key, int value) {
        if (h_.count(key)) {
            Node* node = h_[key];
            node->val = value;
            RemoveNode(node);
            AddNode(node);
        } else {
            if (capacity_ == size_) {
                Node* node = tail_->pre;
                h_.erase(node->key);
                RemoveNode(node);
                // 防止内存泄漏
                delete node;
				size_--;

            }
            size_++;
            Node* node = new Node(key, value);
            AddNode(node);
            h_[key] = node;
            return;
        }
    }

    void AddNode(Node* node) {
        node->next = head_->next;
        node->pre = head_;
        head_->next->pre = node;
        head_->next = node;
    }

    void RemoveNode(Node* node) {
        node->pre->next = node->next;
        node->next->pre = node->pre;
    }

private:
    int capacity_ = 0;
    int size_= 0;
    std::unordered_map<int, Node*> h_;
    Node* head_, *tail_;
};

/**
 * Your LRUCache object will be instantiated and called as such:
 * LRUCache* obj = new LRUCache(capacity);
 * int param_1 = obj->get(key);
 * obj->put(key,value);
*/
