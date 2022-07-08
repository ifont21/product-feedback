import { HTTP_INTERCEPTORS } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AuthHttpInterceptor, AuthModule } from '@auth0/auth0-angular';
import { environment } from '../environments/environment';

import { AppComponent } from './app.component';
import { AuthButtonModule } from './auth-button.component';
import { ProjectModule } from './projects.component';
import { UiHeaderModule } from '@pfeedback/ui-shell';

@NgModule({
  declarations: [AppComponent],
  imports: [
    BrowserModule,
    AuthButtonModule,
    ProjectModule,
    UiHeaderModule,
    AuthModule.forRoot({
      ...environment.auth,
      httpInterceptor: {
        allowedList: [`${environment.dev.serverUrl}/api/*`],
      },
    }),
  ],
  providers: [
    {
      provide: HTTP_INTERCEPTORS,
      useClass: AuthHttpInterceptor,
      multi: true,
    },
  ],
  bootstrap: [AppComponent],
})
export class AppModule {}
