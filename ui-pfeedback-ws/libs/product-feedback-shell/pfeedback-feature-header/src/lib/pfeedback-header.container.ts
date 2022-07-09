import { CommonModule } from '@angular/common';
import { Component, NgModule } from '@angular/core';
import { AuthService } from '@auth0/auth0-angular';
import { PFeedbackUIHeaderModule } from './pfeedback-ui-header.component';

@Component({
  selector: 'app-pfeedback-header-container',
  template: ` <app-pfeedback-ui-header
    (signIn)="authService.loginWithRedirect()"
    (signOut)="authService.logout()"
    [isAuthenticated]="!!(authService.isAuthenticated$ | async)"
    [userProfile]="authService.user$ | async"
  ></app-pfeedback-ui-header>`,
})
export class PFeedbackHeaderContainerComponent {
  constructor(public authService: AuthService) {}
}

@NgModule({
  imports: [CommonModule, PFeedbackUIHeaderModule],
  declarations: [PFeedbackHeaderContainerComponent],
  exports: [PFeedbackHeaderContainerComponent],
})
export class PFeedbackHeaderContainerModule {}
