package src


type Tuple[t1, t2 any] struct {
	First t1
	Second t2
}

type Bestiary struct {
	kills []Tuple[string, int32]
	kills_lookup map[string]int32
	sightings []string
	chats []string
}

func NewBestiary(
	kills []Tuple[string, int32],
	kills_lookup map[string]int32,
	sightings []string,
	chats []string,
) *Bestiary {

	for i, _ := range kills {

		kills_lookup[kills[i].First] = kills[i].Second
	}

	return &Bestiary{
		kills: kills,
		kills_lookup: kills_lookup,
		sightings: sightings,
		chats: chats,
	}
}

func (b *Bestiary) get_kills(entity string) int32 {

	return b.kills_lookup[entity]
}

func (b *Bestiary) add_kills(entity string, count int32) {

	b.kills_lookup[entity] = count

	for i, _ := range b.kills {

		if b.kills[i].First == entity {
			b.kills[i].Second = count
		}
	}

	b.kills = append(b.kills, Tuple[string, int32]{entity, count})

}
