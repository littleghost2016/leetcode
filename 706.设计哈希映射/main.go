package main

import "container/list"

const baseNumber = 769

type entry struct {
	key   int
	value int
}

type MyHashMap struct {
	data []list.List
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{make([]list.List, baseNumber)}
}

func (this *MyHashMap) hash(key int) int {
	return key % baseNumber
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	h := this.hash(key)
	for i := this.data[h].Front(); i != nil; i = i.Next() {
		if tempEntry := i.Value.(entry); tempEntry.key == key {
			i.Value = entry{key, value}
			return
		}
	}
	this.data[h].PushBack(entry{key, value})
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	h := this.hash(key)
	for i := this.data[h].Front(); i != nil; i = i.Next() {
		if tempEntry := i.Value.(entry); tempEntry.key == key {
			return tempEntry.value
		}
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	h := this.hash(key)
	for i := this.data[h].Front(); i != nil; i = i.Next() {
		if tempEntry := i.Value.(entry); tempEntry.key == key {
			this.data[h].Remove(i)
			return
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
