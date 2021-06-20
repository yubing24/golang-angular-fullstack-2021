export interface Language {
  name: string;
}

export interface Country {
  name: string;
  alpha2Code: string;
  alpha3Code: string;
  flagImageUrl: string;
  languages: Language[];
  population: number;
  region: string;
  subRegion: string;
}

export interface SearchCountryStatistics {
  TotalCountries: number;
  RegionCounters: object;
  SubRegionCounters: object;
}
