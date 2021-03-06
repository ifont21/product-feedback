import { CommonModule } from '@angular/common';
import {
  Component,
  EventEmitter,
  Input,
  NgModule,
  Output,
} from '@angular/core';
import { User } from '@auth0/auth0-angular';

@Component({
  selector: 'app-pfeedback-ui-header',
  template: `<div class="h-14 px-6 py-8">
    <div class="flex justify-between">
      <div>
        <span>PFeedback</span>
      </div>
      <ul class="flex items-center">
        <li>
          <button
            class="bg-indigo-900 text-neutral-50 px-6 py-2 rounded-lg"
            (click)="signIn.emit()"
            *ngIf="isAuthenticated === false"
          >
            Sign In
          </button>
          <span *ngIf="isAuthenticated">
            {{ userProfile?.name }}
          </span>
        </li>
        <li *ngIf="isAuthenticated">
          <button
            class="text-neutral-900 px-6 py-2 rounded-lg"
            (click)="signOut.emit()"
          >
            Sign Out
          </button>
        </li>
      </ul>
    </div>
  </div> `,
})
export class PFeedbackUIHeaderComponent {
  @Output()
  signIn: EventEmitter<void> = new EventEmitter<void>();

  @Output()
  signOut: EventEmitter<void> = new EventEmitter<void>();

  @Input()
  isAuthenticated!: boolean | null;

  @Input()
  userProfile: User | undefined | null;
}

@NgModule({
  imports: [CommonModule],
  declarations: [PFeedbackUIHeaderComponent],
  exports: [PFeedbackUIHeaderComponent],
})
export class PFeedbackUIHeaderModule {}
