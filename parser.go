package log2ltsv

type Parser interface {
	ParseAndOutput(string) (string, error)
}
