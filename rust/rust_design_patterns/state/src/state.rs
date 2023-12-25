use cursive::views::TextView;

use crate::player::Player;

pub struct StoppedState;
pub struct PausedState;
pub struct PlayingState;

pub trait State {
    fn play(self: Box<Self>, player: &mut Player) -> Box<dyn State>;
    fn stop(self: Box<Self>, player: &mut Player) -> Box<dyn State>;
    fn render(&self, player: &Player, view: &mut TextView);
}

impl State for StoppedState {
    fn play(self: Box<Self>, player: &mut Player) -> Box<dyn State> {
        player.play();

        Box::new(PlayingState)
    }

    fn stop(self: Box<Self>, _: &mut Player) -> Box<dyn State> {
        self
    }

    fn render(&self, _: &Player, view: &mut TextView) {
        view.set_content("[Stopped] press 'Play'")
    }
}

impl State for PausedState {
    fn play(self: Box<Self>, player: &mut Player) -> Box<dyn State> {
        player.pause();

        Box::new(PlayingState)
    }

    fn stop(self: Box<Self>, player: &mut Player) -> Box<dyn State> {
        player.pause();
        player.rewind();

        Box::new(StoppedState)
    }

    fn render(&self, player: &Player, view: &mut TextView) {
        view.set_content(format!(
            "[Paused] {} - {} sec",
            player.track().title,
            player.track().duration
        ))
    }
}

impl State for PlayingState {
    fn play(self: Box<Self>, player: &mut Player) -> Box<dyn State> {
        player.pause();

        Box::new(PausedState)
    }

    fn stop(self: Box<Self>, player: &mut Player) -> Box<dyn State> {
        player.pause();
        player.rewind();

        Box::new(StoppedState)
    }

    fn render(&self, player: &Player, view: &mut TextView) {
        view.set_content(format!(
            "[Playing] {} - {} sec",
            player.track().title,
            player.track().duration
        ))
    }
}

impl dyn State {
    pub fn next(self: Box<Self>, player: &mut Player) -> Box<dyn State> {
        player.next_track();

        self
    }

    pub fn prev(self: Box<Self>, player: &mut Player) -> Box<dyn State> {
        player.prev_track();

        self
    }
}
