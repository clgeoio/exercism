//Package tree building allows generation of a generic parent/child tree
package tree

import (
	"errors"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

//Build generates a Node tree based on input records.
func Build(records []Record) (*Node, error) {

	recordLength := len(records)
	if recordLength == 0 {
		return nil, nil
	}

	lastID := 0
	tree := Node{ID: 0}
	sortedRecords := map[int]Record{}
	sortedMaxId := 0
	idPointers := map[int]*Node{}
	idPointers[0] = &tree

	//sort the records based on ID, so we can process top down, worst case O(n)
	for _, r := range records {
		sortedRecords[r.ID] = r
		//keep note of the maximum ID for error checking if the ID is out of bounds
		if r.ID > sortedMaxId {
			sortedMaxId = r.ID
		}
	}

	if sortedMaxId != (recordLength - 1) {
		return nil, errors.New("No root record or ID exceeds record length")
	}

	//loop through records worst case O(n)
	for i := 0; i < recordLength; i++ {
		rec := sortedRecords[i]
		//special treatment for root record
		if i == 0 {
			if !(rec.ID == 0 && rec.Parent == 0) {
				return nil, errors.New("Invalid root record")
			}
		} else {
			if rec.ID != (lastID + 1) {
				return nil, errors.New("Current record ID not greater than last ID")
			}

			if i != 0 && rec.ID <= rec.Parent {
				return nil, errors.New("Current record ID equal to or less than Parent")
			}

			parent, ok := idPointers[rec.Parent]
			if !ok {
				return nil, errors.New("Could not find parent Node")
			}
			child := Node{ID: rec.ID}
			parent.Children = append(parent.Children, &child)
			idPointers[rec.ID] = &child
		}

		lastID = rec.ID
	}
	return &tree, nil
}
