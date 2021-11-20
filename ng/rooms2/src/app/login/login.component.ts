import { Component, OnInit, Output, EventEmitter } from '@angular/core';
import { FormControl } from '@angular/forms';

import { TokenService } from "../services/token.service";
import { TokenModel } from "../models/token.model";

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})

export class LoginComponent implements OnInit {
  isLoggedIn = false;
  token = "";
  username = new FormControl('');
  password = new FormControl('');

  @Output() messageEvent = new EventEmitter<string>();
  
  constructor(
    private tokenService: TokenService
    ) { 
      if (localStorage.getItem('token') != null && localStorage.getItem('token') != "") {
        this.token = localStorage.getItem('token') || "";
        this.isLoggedIn = true;
        this.username.setValue(localStorage.getItem('username'));
      }
    }

  ngOnInit(): void { }

  onLogout() : void {
    console.log("Logging out user: " + this.username.value);

    this.token = "";
    this.isLoggedIn = false;

    this.updateLocalStorage();

    this.messageEvent.emit("User '" + this.username.value + "' successfully logged out.");
  }

  onLogin(): void {
    console.log("Logging in user: " + this.username.value);

    this.tokenService.getAuthToken(this.username.value, this.password.value)
    .subscribe ( 
      (data : TokenModel) => this.handleAuthSuccess(data),
      (error) => this.handleAuthFailure(error)
    )
  }

  private handleAuthSuccess(data : TokenModel) {
    console.log("Login success.");
  
    this.token = data.token;
    this.isLoggedIn = true;
    this.updateLocalStorage();

    this.messageEvent.emit("User '" + this.username.value + "' successfully logged in.");
  }

  private handleAuthFailure(error: any) {
    console.log("Login failed.");
    console.log(error);

    this.token = "";
    this.isLoggedIn = false;
    this.updateLocalStorage();

    this.messageEvent.emit("Login failed for user '" + this.username.value + "'");
  }

  private updateLocalStorage(){
    localStorage.setItem('token', this.token);
    localStorage.setItem('username', this.username.value);
  }
}
