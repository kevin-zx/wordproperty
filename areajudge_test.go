package wordproperty

import (
	"fmt"
	"testing"
)

func TestIsChinaArea(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 float64
		//want2 Area
	}{
		{
			name:  "1",
			args:  args{
				key: "上海",
			},
			want:  true,
			want1: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := IsChinaArea(tt.args.key)
			if got != tt.want {
				t.Errorf("IsChinaArea() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("IsChinaArea() got1 = %v, want %v", got1, tt.want1)
			}
			fmt.Printf("%+v\n", got2)
			//if !reflect.DeepEqual(got2, tt.want2) {
			//	t.Errorf("IsChinaArea() got2 = %v, want %v", got2, tt.want2)
			//}
		})
	}
}
