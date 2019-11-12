package sqlsmith

import "log"

func (s *SQLSmith) debugPrint(a ...interface{}) {
	if s.debug {
		log.Println(a...)
	}
}
