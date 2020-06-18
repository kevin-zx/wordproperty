package wordproperty

import "testing"

func TestEnvWordProperty(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name       string
		args       args
		isEnv      bool
		confidence float64
	}{
		{
			name:       "补水喷雾的正确使用方法",
			args:       args{
				word: "补水喷雾的正确使用方法",
			},
			isEnv:      true,
			confidence: 0.2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := EnvWordProperty(tt.args.word)
			if got != tt.isEnv {
				t.Errorf("EnvWordProperty() got = %v, isEnv %v", got, tt.isEnv)
			}
			if got1 != tt.confidence {
				t.Errorf("EnvWordProperty() got1 = %v, isEnv %v", got1, tt.confidence)
			}
		})
	}
}
