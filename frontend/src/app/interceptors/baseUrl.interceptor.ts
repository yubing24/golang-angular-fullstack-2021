import {Injectable} from '@angular/core';
import {HttpEvent, HttpHandler, HttpInterceptor, HttpRequest} from '@angular/common/http';
import {Observable} from 'rxjs';
import {environment} from '@env/environment';

@Injectable()
export class BaseUrlInterceptor implements HttpInterceptor {
  baseUrl = environment.apiBaseUrl;
  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    const apiReq = req.clone({ url: `${this.baseUrl}${req.url}`});
    return next.handle(apiReq);
  }
}
