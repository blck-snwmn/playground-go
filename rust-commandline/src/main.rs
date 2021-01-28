use clap::Clap;

// can't rename `Option`
#[derive(Clap)]
#[clap(version = "1.0", author = "snowan")]
struct Opts {
    // hello!
    #[clap(short, long, default_value = "hello world!")]
    greet_message: String,
    // anything
    anything: String,
    // subcommand
    #[clap(subcommand)]
    subcmd: SubCommand,
}

#[derive(Clap)]
enum SubCommand {
    #[clap(version = "1.3", author = "snowman")]
    Run(Mode),
}

#[derive(Clap)]
struct Mode {
    // if true, dry run
    #[clap(long)]
    dry_run: bool,
}

fn main() {
    let opts: Opts = Opts::parse();

    println!("anything: {}", opts.anything);
    println!("hello! {}", opts.greet_message);

    match opts.subcmd {
        SubCommand::Run(t) => {
            if t.dry_run {
                println!("rollback");
            } else {
                println!("commit");
            }
        }
    }
}
