import { Component, Input, NgModule } from '@angular/core';
import { BoardDetailDTO } from '@pfeedback/board-data-spec';

@Component({
  selector: 'app-board-feedbacks-heading',
  template: `<div
    class="h-36 rounded-lg flex items-end p-6 bg-[url('assets/background-header.png')] bg-no-repeat bg-cover"
  >
    <div>
      <h1 class="text-xl text-white font-bold tracking-tight">{{ boardHeading?.project }}</h1>
      <h2 class="text-sm text-white/75 font-medium ">{{ boardHeading?.name }} Board</h2>
    </div>
  </div>`,
})
export class BoardFeedbackHeadingComponent {
  @Input()
  boardHeading: BoardDetailDTO | undefined | null;
}

@NgModule({
  imports: [],
  declarations: [BoardFeedbackHeadingComponent],
  exports: [BoardFeedbackHeadingComponent],
})
export class BoardFeedbackHeadingModule {}
