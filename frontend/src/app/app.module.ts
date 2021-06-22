import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { SearchCountryComponent } from '@components/search-country/search-country.component';
import {HttpClientModule} from '@angular/common/http';
import {SearchService} from '@services/search.service';
import {INTERCEPTOR_PROVIDERS} from '@interceptors/interceptor-providers';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import { SearchCountryStatisticsComponent } from './components/search-country-statistics/search-country-statistics.component';

@NgModule({
  declarations: [
    AppComponent,
    SearchCountryComponent,
    SearchCountryStatisticsComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
  ],
  providers: [
    INTERCEPTOR_PROVIDERS,
    SearchService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
