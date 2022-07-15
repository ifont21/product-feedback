import { Component, Input, NgModule } from '@angular/core';

@Component({
  selector: 'app-ui-vote',
  template: `
    <div class="bg-zinc-100 rounded-lg py-2 w-10 flex flex-col items-center space-y-2">
      <div>
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="16"
          height="16"
          fill="currentColor"
          class="bi bi-chevron-up fill-blue-700"
          viewBox="0 0 16 16"
        >
          <path
            fill-rule="evenodd"
            d="M7.646 4.646a.5.5 0 0 1 .708 0l6 6a.5.5 0 0 1-.708.708L8 5.707l-5.646 5.647a.5.5 0 0 1-.708-.708l6-6z"
          />
        </svg>
      </div>
      <span class="text-sm tracking-tight font-bold text-indigo-900">{{
        votes
      }}</span>
    </div>
  `,
})
export class UIVoteComponent {
  @Input()
  votes = 0;
}

@NgModule({
  declarations: [UIVoteComponent],
  exports: [UIVoteComponent],
})
export class UIVoteModule {}
