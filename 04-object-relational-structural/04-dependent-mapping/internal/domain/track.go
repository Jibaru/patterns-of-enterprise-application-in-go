package domain

// Track represents a track in an album.
type Track struct {
	ID       int    // Identity field for the track
	Title    string // Title of the track
	Duration int    // Duration of the track in seconds
}
