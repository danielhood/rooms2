import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';

import { CommandResponseModel } from '../models/commandresponse.model';

@Injectable({
  providedIn: 'root',
})
export class CommandService {
  constructor(private http: HttpClient) { }

  buildRequestOptions(): object {
    console.log("Setting token: " + localStorage.getItem('token'))
    return {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + localStorage.getItem('token'),
        'Access-Control-Allow-Origin': '*',
      })
    };
  }

  sendCommand(command: string): Observable<CommandResponseModel> {
    var username = localStorage.getItem('username');
    console.log('Sending command for user ' + username);
    return this.http.get<CommandResponseModel>('https://rooms2.local:8443/command?u=' + username + '&c=' + encodeURIComponent(command), 
      this.buildRequestOptions())
  } 
}