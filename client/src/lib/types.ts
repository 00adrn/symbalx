export class navButton {
    text: string = '';
    redirect: string = '';
}

export interface imageItem {
    url: string;
    height: number;
    width: number;
}

class spotifyItem {
    id: string = '';
    name: string = '';
    uri: string = '';
}

export class artistItem extends spotifyItem { }

export class albumItem extends spotifyItem{
    total_tracks: number = 0;
    artists: artistItem[] = [];
    images: imageItem[] = [];

    get getImage(): string { return this.images.length > 0 ? this.images[0].url : ''; }
}

export class trackItem extends spotifyItem {
    album: albumItem = new albumItem();
    artists: artistItem[] = [];

    get getImage(): string { return this.album.getImage; }
}