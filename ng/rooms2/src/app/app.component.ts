import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  currMessage = "Welcome to rooms2!";

  receiveMessage(event: string) {
    this.currMessage = event
  }
}
