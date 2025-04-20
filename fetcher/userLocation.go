package fetcher

import "path"

func (ul UserLocation) UserId() string {
	return path.Base(string(ul))
}
