
// -*- JavaScript -*-


class MemberItem extends React.Component {
  render() {
    return (
      <tr>
        <td> {this.props.id}    </td>
        <td> {this.props.firstname} </td>
        <td> {this.props.lastname}  </td>
      </tr>
    );
  }
}

class MembersList extends React.Component {
  constructor(props) {
    super(props);
    this.state = { people: [] };
  }

  componentDidMount() {
    this.serverRequest =
      axios
        .get("/members")
        .then((result) => {
           this.setState({ members: result.data });
        });
  }

  render() {
    const members = this.state.members.map((member, i) => {
      return (
        <MemberItem key={i} id={member.Id} first={member.Firstname} last={member.Lastname} />
      );
    });

    return (
      <div>
        <table><tbody>
          <tr><th>Id</th><th>Firstname</th><th>Lastname</th></tr>
          {members}
        </tbody></table>

      </div>
    );
  }
}

ReactDOM.render( <MembersList/>, document.querySelector("#root"));
