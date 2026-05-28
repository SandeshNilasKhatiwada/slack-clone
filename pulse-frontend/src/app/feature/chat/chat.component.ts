import { Component, inject, OnInit, signal, OnDestroy } from '@angular/core';
import { ReactiveFormsModule, FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ChatService } from '../../core/services/chat.service';


@Component({
  selector: 'app-chat',
  standalone: true,
  imports: [ReactiveFormsModule],
  templateUrl: './chat.component.html',
})
export class ChatComponent implements OnInit, OnDestroy {
  private chatService = inject(ChatService);
  private fb = inject(FormBuilder);

  // Angular Signal to hold our array of messages
  messages = signal<{ text: string }[]>([]);

  chatForm: FormGroup = this.fb.group({
    messageText: ['', Validators.required]
  });

  ngOnInit() {
    this.chatService.connect();

    // Listen for incoming messages and update the Signal
    this.chatService.socket$.subscribe({
      next: (msg: any) => {
        // Update the signal by appending the new message
        this.messages.update(currentMessages => [...currentMessages, msg]);
      }
    });
  }

  sendMessage() {
    if (this.chatForm.valid) {
      const text = this.chatForm.value.messageText;
      this.chatService.sendMessage(text);

      // Clear the input box after sending
      this.chatForm.reset();
    }
  }

  ngOnDestroy() {
    // Cleanup if the user leaves the chat page
    if (this.chatService.socket$) {
      this.chatService.socket$.complete();
    }
  }
}
