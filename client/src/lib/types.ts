export class navButton {
    text: string = '';
    redirect: string = '';
}

export class imageItem {
    url: string = '';
    height: number = 0;
    width: number = 0;

    constructor (data?: any) {
        if (!data) return;

        this.url = data.url;
        this.height = data.height;
        this.width = data.width;
    }
}

class spotifyItem {
    id: string = '';
    name: string = '';
    uri: string = '';

    constructor (data?: any) {
        if (!data) return;

        this.id = data.artists
        this.name = data.name
        this.uri = data.uri
    }
}

export class artistItem extends spotifyItem {
    
}

export class albumItem extends spotifyItem{
    total_tracks: number = 0;
    artists: artistItem[] = [];
    images: imageItem[] = [];

    constructor (data?: any) {
        super(data);
        if (!data) return;

        this.total_tracks = data.total_tracks;
        this.artists = data.artists.map((artist: any) => new artistItem(artist));
        this.images = data.images.map((image: any) => new imageItem(image))
    }

    get getImage(): string { 
        return this.images.length === 0 ? '' : this.images[0].url; 
    }
}

export class trackItem extends spotifyItem {
    album: albumItem = new albumItem();
    artists: artistItem[] = [];
    images: imageItem[] = [];

    constructor (data?: any) {
        super(data);
        if (!data) return;

        this.album = new albumItem(data.album);
        this.artists = data.artists.map((artist: any) => new artistItem(artist));
        this.images = data.album.images.map((image: any) => new imageItem(image))
    }

    get getImage(): string { 
        return this.album.getImage; 
    }
}

export class spotifyUserItem extends spotifyItem {
    display_name: string = '';
    images: imageItem[] = [];

    constructor (data: any) {
        super(data);
        
        this.display_name = data.display_name;
        this.images = data.images.map((image: any) => new imageItem(image))
    }

    get getImage(): string {
        return this.images.length === 0 ? '' : this.images[0].url;
    }
}