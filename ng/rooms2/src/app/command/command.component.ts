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

  @Output() messageEvent = new EventEmitter<string>();

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

    this.command.setValue("");
  }

  handleCommandFailure(error: any) {
    console.log("Send command failed");
    console.log(error);

    this.messageEvent.emit("Error processing command.");
  }

}
