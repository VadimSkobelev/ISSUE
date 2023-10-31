package postgres

import (
	"log"
	"os"
	"testing"
)

var s *Storage

func TestMain(m *testing.M) {
	constr := "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
	var err error
	s, err = New(constr)
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(m.Run())
}

func TestStorage_Tasks(t *testing.T) {
	type args struct {
		taskID   int
		authorID int
	}
	tests := []struct {
		name    string
		s       *Storage
		args    args
		want    []Task
		wantErr bool
	}{
		// TODO: Add test cases.
		{"All tasks", s, args{0, 0}, []Task{}, false},
		{"First task", s, args{1, 0}, []Task{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Tasks(tt.args.taskID, tt.args.authorID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Tasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("Storage.Tasks() = %v, want %v", got, tt.want)
			// }
			t.Log(got)
		})
	}
}

func TestStorage_NewTask(t *testing.T) {
	type args struct {
		t Task
	}
	tests := []struct {
		name    string
		s       *Storage
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Create task", s, args{Task{Title: "Unit Test", Content: "Testing"}}, 19, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.NewTask(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.NewTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.NewTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_TaskByLabel(t *testing.T) {
	type args struct {
		label string
	}
	tests := []struct {
		name    string
		s       *Storage
		args    args
		want    []Task
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Select by lable", s, args{"first"}, []Task{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.TaskByLabel(tt.args.label)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.TaskByLabel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("Storage.TaskByLabel() = %v, want %v", got, tt.want)
			// }
			t.Log(got)
		})
	}
}

func TestStorage_UpdateTask(t *testing.T) {
	type args struct {
		t Task
	}
	tests := []struct {
		name    string
		s       *Storage
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Update task with id=4", s, args{Task{ID: 4, Title: "Final-16", Content: "tests!!!-16"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.UpdateTask(tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("Storage.UpdateTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_DeleteTask(t *testing.T) {
	type args struct {
		t Task
	}
	tests := []struct {
		name    string
		s       *Storage
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Delete task with id=11", s, args{Task{ID: 11}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.DeleteTask(tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("Storage.DeleteTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
