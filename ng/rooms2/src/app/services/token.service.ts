import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpErrorResponse } from '@angular/common/http';
import { Observable, throwError, catchError } from 'rxjs';

import { TokenModel } from '../models/token.model';

@Injectable({
  providedIn: 'root',
})
export class TokenService {
  constructor(private http: HttpClient) { }

  buildTokenRequestBody(user: string, pass: string): object {
    return {
      username: user,
      password: pass
    }
  }

  buildTokenRequestOptions(): object {
    return {
      headers: new HttpHeaders({
        'Access-Control-Allow-Origin': '*',
      })
    };
  }

  getAuthToken(user: string, pass: string): Observable<TokenModel> {
    console.log('Geting auth token...');
    return this.http.post<TokenModel>('https://rooms2.local:8443/token', 
      this.buildTokenRequestBody(user, pass), 
      this.buildTokenRequestOptions())
  } 

  storeToken(token: string) {
    localStorage.setItem('token', token);
  }

}