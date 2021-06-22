import {Component, OnInit} from '@angular/core';
import {SearchService} from '@services/search.service';
import {Country, SearchCountryStatistics} from '@app/models/country';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';

@Component({
  selector: 'app-search-country',
  templateUrl: './search-country.component.html',
  styleUrls: ['./search-country.component.scss']
})
export class SearchCountryComponent implements OnInit {

  public searchButtonDisabled = false;

  public errorMessage: string | null = null;

  searchForm: FormGroup = this.formBuilder.group({
    name: ['', [Validators.required]]
  });


  get name(): any {
    return this.searchForm.get('name');
  }

  public results: any[] = [];
  public statistics: any = null;

  constructor(
    private searchService: SearchService,
    private formBuilder: FormBuilder
  ) {
  }

  get hasResults(): boolean {
    return !!this.results && this.results.length > 0;
  }

  ngOnInit(): void {
  }

  public searchCountries(evt: any): void {
    evt.preventDefault();
    this.searchButtonDisabled = true;
    this.errorMessage = null;
    const form = {
      name: this.name.value.toString(),
      code: this.name.value.toString()
    };
    this.results = [];
    this.searchService.searchCountry(form).subscribe(res => {
      this.results = res.data.records as Country[];
      this.statistics = res.data.statistics as SearchCountryStatistics;
      this.searchButtonDisabled = false;
    }, err => {
      switch (err.error.status) {
        case 404:
          this.errorMessage = `Error: ${err.error.message}`;
          break;
        default:
          alert(`Error: ${JSON.stringify(err.error.message)}`);
      }
      this.searchButtonDisabled = false;
    });
  }
}
