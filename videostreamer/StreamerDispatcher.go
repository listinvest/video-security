package videostreamer

//StreamerDispatcher contains rtsp streams
type StreamerDispatcher struct {
	Streamers []IStreamer
}

//Get item
func (d *StreamerDispatcher) Get(key string) IStreamer {
	if d.Streamers == nil {
		return nil
	}

	index := d.indexOf(key)
	return d.Streamers[index]
}

//Add item
func (d *StreamerDispatcher) Add(s IStreamer) {
	if d.Streamers == nil {
		d.Streamers = make([]IStreamer, 0)
	}

	d.Streamers = append(d.Streamers, s)
}

//Remove item
func (d *StreamerDispatcher) Remove(key string) bool {
	index := d.indexOf(key)
	if index > -1 {
		d.Streamers = d.removeByIndex(index)
	}
	return false
}

//indexOf
func (d *StreamerDispatcher) indexOf(key string) int {
	if d.Streamers == nil {
		return -1
	}

	for index, s := range d.Streamers {
		if s.GetKey() == key {
			return index
		}
	}
	return -1
}

//removeByIndex
func (d *StreamerDispatcher) removeByIndex(i int) []IStreamer {
	d.Streamers[len(d.Streamers)-1], d.Streamers[i] = d.Streamers[i], d.Streamers[len(d.Streamers)-1]
	return d.Streamers[:len(d.Streamers)-1]
}
