import {Injectable} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {SearchCountryForm} from '../forms/search-country';
import {Observable} from 'rxjs';
import {ServerResponse} from './response';

@Injectable()
export class SearchService {
  constructor(
    private http: HttpClient
  ) {}

  searchCountry(form: SearchCountryForm): Observable<ServerResponse> {
    const param = new HttpParams().set('name', form.name).set('code', form.name);
    return this.http.get<ServerResponse>('/search-country', {
      params: param
    });
  }
}
