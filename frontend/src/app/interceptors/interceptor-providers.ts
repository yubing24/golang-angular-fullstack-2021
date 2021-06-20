import {HTTP_INTERCEPTORS} from '@angular/common/http';
import {BaseUrlInterceptor} from '@interceptors/baseUrl.interceptor';

export const INTERCEPTOR_PROVIDERS = [
  {
    provide: HTTP_INTERCEPTORS,
    useClass: BaseUrlInterceptor,
    multi: true
  }
];
