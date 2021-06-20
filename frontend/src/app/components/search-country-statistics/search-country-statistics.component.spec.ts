import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SearchCountryStatisticsComponent } from './search-country-statistics.component';

describe('SearchCountryStatisticsComponent', () => {
  let component: SearchCountryStatisticsComponent;
  let fixture: ComponentFixture<SearchCountryStatisticsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SearchCountryStatisticsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(SearchCountryStatisticsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
