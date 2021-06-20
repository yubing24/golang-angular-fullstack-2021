import { Component, OnInit } from '@angular/core';
import {SearchService} from '../../services/search.service';
import {Country, SearchCountryStatistics} from '@app/models/country';

@Component({
  selector: 'app-search-country',
  templateUrl: './search-country.component.html',
  styleUrls: ['./search-country.component.scss']
})
export class SearchCountryComponent implements OnInit {

  public searchCriteria  = '';
  public searchButtonDisabled = false;

  constructor(
    private searchService: SearchService
  ) { }

  get hasResults(): boolean {
    return !!this.results && this.results.length > 0;
  }

  public results: any[] = [];
  public statistics: any = null;

  ngOnInit(): void {
  }

  public searchCountries(): void {
    this.searchButtonDisabled = true;
    const form = {
      name: this.searchCriteria,
      code: this.searchCriteria
    };
    this.results = [];
    this.searchService.searchCountry(form).subscribe(res => {
      this.results = res.data.records as Country[];
      this.statistics = res.data.statistics as SearchCountryStatistics;
      this.searchButtonDisabled = false;
    }, err => {
      console.log(`${JSON.stringify(err)}`);
      alert(`Error: ${JSON.stringify(err.error.message)}`);
      this.searchButtonDisabled = false;
    });
  }

  public clear(): void {
    this.searchCriteria = '';
    this.results = [];
  }
}
