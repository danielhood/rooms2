import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { TokenService } from "../services/token.service";
import { TokenModel } from "../models/token.model";
import { MessageService } from "../services/message.service";

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})

export class LoginComponent implements OnInit {
  isLoggedIn = false;
  token = "";
  loggedInUsername = "";
  
  username = new FormControl('');
  password = new FormControl('');

  constructor(
    private tokenService: TokenService,
    private messageService: MessageService
    ) { 
      if (localStorage.getItem('token') != null && localStorage.getItem('token') != "") {
        this.setLoggedInUser(localStorage.getItem('username') || "", localStorage.getItem('token') || "")
      }
    }

  ngOnInit(): void {
    if (localStorage.getItem('token') != null && localStorage.getItem('token') != '') {
      this.currentLoginSuccessMessage();     
    }
  }

  onLogout() : void {
    console.log("Logging out user: " + this.username.value);

    this.clearLoggedInUser();

    this.messageService.setMessage("Successfully logged out.");
  }

  onLogin(): void {
    console.log("Logging in user: " + this.username.value);
    this.messageService.setMessage("Logging in user: " + this.username.value + "...");

    this.tokenService.getAuthToken(this.username.value, this.password.value)
    .subscribe ( 
      (data : TokenModel) => this.handleAuthSuccess(data),
      (error) => this.handleAuthFailure(error)
    )
  }

  private currentLoginSuccessMessage(){ 
    this.messageService.setMessage("User '" + this.loggedInUsername + "' successfully logged in.");
  }

  private handleAuthSuccess(data : TokenModel) {
    console.log("Login success.");
  
    this.setLoggedInUser(this.username.value, data.token);

    this.currentLoginSuccessMessage();   

    // Clear values
    this.username.setValue("");
    this.password.setValue("");
  }

  private handleAuthFailure(error: any) {
    console.log("Login failed.");
    console.log(error);

    this.clearLoggedInUser();

    this.messageService.setMessage("Login failed for user '" + this.username.value + "'");
  }

  private setLoggedInUser(username: string, token: string) {
    this.loggedInUsername = username;
    this.token = token;
    this.isLoggedIn = true;
    this.updateLocalStorage(); 
  }

  private clearLoggedInUser() {
    this.loggedInUsername = "";
    this.token = "";
    this.isLoggedIn = false;
    this.updateLocalStorage();
  }

  private updateLocalStorage(){
    localStorage.setItem('token', this.token);
    localStorage.setItem('username', this.loggedInUsername);
  }
}
