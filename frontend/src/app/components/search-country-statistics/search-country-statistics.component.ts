import {Component, Input, OnInit} from '@angular/core';
import {SearchCountryStatistics} from '@app/models/country';

@Component({
  selector: 'app-search-country-statistics',
  templateUrl: './search-country-statistics.component.html',
  styleUrls: ['./search-country-statistics.component.scss']
})
export class SearchCountryStatisticsComponent implements OnInit {
  @Input() statistics: SearchCountryStatistics | null = null;
  constructor() { }
  ngOnInit(): void {
  }
}
