import { Component, OnInit, Output, EventEmitter } from '@angular/core';
import { FormControl } from '@angular/forms';

import { CommandService } from "../services/command.service";
import { CommandResponseModel } from "../models/commandresponse.model";
import { MessageService } from "../services/message.service";

@Component({
  selector: 'app-command',
  templateUrl: './command.component.html',
  styleUrls: ['./command.component.scss']
})
export class CommandComponent implements OnInit {

  command = new FormControl('');

  constructor(
    private commandService: CommandService,
    private messageService: MessageService
    ) { }

  ngOnInit(): void {
  }

  onCommandEnter(): void {
    console.log("Sending command: " + this.command.value);

    this.commandService.sendCommand(this.command.value)
    .subscribe ( 
      (data : CommandResponseModel) => this.handleCommandResponse(data),
      (error) => this.handleCommandFailure(error)
    )
  }

  handleCommandResponse(commandResponse: CommandResponseModel) {
    console.log("Received command response:" );
    console.log(commandResponse);

    this.messageService.setMessages(commandResponse.responses);

    this.messageService.setMessages(
      ["This is a long description of a room that has many special things about it that all need to be described in long winded, Stephen King style, description more than any other aspect of the adventure such that it will immerse the user even further into the enviroment and cause them to cherish this story for ever and ever and ever, until the end of your days, wherever you may end up.",
    "You are aslo on fire.",
    "And grappled by a troll."]
    );

    this.command.setValue("");
  }

  handleCommandFailure(error: any) {
    console.log("Send command failed");
    console.log(error);

    this.messageService.setMessage("Unable to send command to server.");
  }

}
